package main

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
	"moderatorBot/internal/storage"
	"strconv"
	"sync"
	"time"
)

type (
	// BaseChat  база для всех типов чатов
	BaseChat struct {
		channel chan tgbotapi.Update
		db      storage.ChatIdModel
		tg      int64
	}

	// ChatInterface интерфейс чатов
	ChatInterface interface {
		// routine слушает и обрабатывает сообщения определенного чата
		routine(*tgbotapi.BotAPI, map[int64]ChatInterface, *sync.Mutex, storage.Interface)
		// send отправляет чату инвормацию о телеграм-сообщении
		send(tgbotapi.Update)
	}

	// ChatStatus текущее состояние чата. Пожалуй, нахуй не нужны. В PrivateChat будем использовать дочернии обработчики
	ChatStatus byte

	// PrivateChat структура для работы с приватным чатом
	PrivateChat struct{ BaseChat }

	// SupergroupChat структура для работы с общим чатом
	SupergroupChat struct {
		BaseChat
		moderated bool
	}
)

func (bc BaseChat) send(update tgbotapi.Update) { bc.channel <- update }

func (pc PrivateChat) routine(_ *tgbotapi.BotAPI, chats map[int64]ChatInterface, mainMutex *sync.Mutex, storage storage.Interface) {
	lastMassageTime := time.After(time.Hour * 10)
	for {
		select {
		case message := <-pc.channel:
			//close(lastMassageTime)
			lastMassageTime = time.After(time.Hour * 10)
			log.Println("MainStatus.Update")
			if message.Message != nil {
				if message.Message.Chat == nil {
					// TODO: пишем в лог
					continue
				}
				if message.Message.Chat.Type != privateChatType {
					// TODO: пишем в лог
					continue
				}
				switch message.Message.Text {
				case AddAdminText:
					AdminAddition(pc.tg, pc.channel, storage)
				case AddBannedWordText:
					BanWordAddition(pc.tg, pc.channel, storage)
				case SetBanTimeText:
					setBantime(pc.tg, storage, pc.channel)
				case SetWarningsText:
					SetWarningsVal(pc.tg, storage, pc.channel)
				case GetSettingsText:
					GetPanishments(pc.tg, storage)
				case DeleteBannedWordText:
					ContainsPolicy = DeleteBannedWord(pc.tg, pc.channel, storage)
				default:
					showCmd := tgbotapi.NewMessage(pc.tg, WhatToDoText)
					showCmd.ReplyMarkup = MainAdminKeyboard
					BotAPI.Request(showCmd)
				}

			}

		case <-lastMassageTime:
			log.Println("time out")
			mainMutex.Lock()
			log.Println("Chat " + strconv.FormatInt(pc.tg, 10) + " deleted")
			log.Println(chats)
			delete(chats, pc.tg)
			mainMutex.Unlock()
		}
	}
}

func (sc SupergroupChat) routine(botApi *tgbotapi.BotAPI, chats map[int64]ChatInterface, mainMutex *sync.Mutex, storage storage.Interface) {
	lastMassageTime := time.After(time.Hour * 10)
	//status := nilStatus
	for {
		select {
		case message := <-sc.channel:
			switch {
			case message.Message != nil:
				// Проверка целостности входящих данных
				if message.Message.Chat == nil {
					// TODO: пишем в лог
					continue
				}
				if message.Message.Chat.Type != supergroupChatType && message.Message.Chat.Type != groutChatType {
					// TODO: пишем в лог
					continue
				}
				banned := storage.GetBanList()
				for i := 0; i < len(banned); i++ {
					temp, _ := strconv.Atoi(banned[i])
					if message.Message.From.ID == int64(temp) {
						dm := tgbotapi.NewDeleteMessage(message.Message.Chat.ID, message.Message.MessageID)
						if _, fail := botApi.Request(dm); fail != nil {
							// TODO: сохраняем кучу данных в лог
						}
					}
				}
				// Цикл проверок
				log.Println(len(ContainsPolicy), ContainsPolicy)

				for _, v := range ContainsPolicy {
					if err := v.Check(message); err != nil {
						log.Println("ААААА!", err.Error()) // вердикт (.)(.)
						// Если проверка сработала, то удаляем сообщение
						storage.Crime(message.Message.From.ID, panishments.Warnings, panishments.Bandur)
						dm := tgbotapi.NewDeleteMessage(message.Message.Chat.ID, message.Message.MessageID)
						if _, fail := botApi.Request(dm); fail != nil {
							// TODO: тут снова туду сохнарить в базу сообщение об ошибке на третьей стороне (')(')
						}
						break
					}
				}
			default:
				// Пришло что-то, что мы не умеем обрабатывать. Или лень
				// TODO: пишем message в лог
			}
		case <-lastMassageTime:
			// TODO: Доработать: убрать удаление сата из мапы в main
			log.Println("time out")
			mainMutex.Lock()
			log.Println("Chat " + strconv.FormatInt(sc.tg, 10) + " deleted")
			log.Println(chats)
			delete(chats, sc.tg)
			mainMutex.Unlock()
		}
	}
}
