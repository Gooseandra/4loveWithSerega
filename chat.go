package main

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
	"moderatorBot/internal/policy"
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
		policies  []policy.Interface
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
				case "add admin":
					AdminAddition(pc.tg, pc.channel, storage)
				case "ban word":
					if IsItAdmin(pc.tg, storage) == true {
						banWord, err := InputText(pc.tg, pc.channel, "Введите запрещённое слово")
						if err != nil {
							//TODO: какая то реакция
						}
						disc, err := InputText(pc.tg, pc.channel, "NewContains/NewStartWith?")
						if err != nil {
							//TODO: какая то реакция
						}
						storage.AddBannedWord(banWord, disc)
					} else {
						BotAPI.Send(tgbotapi.NewMessage(pc.tg, "Вы не являетесь админом"))
					}
				case "add me":
					BotAPI.Send(tgbotapi.NewMessage(message.Message.Chat.ID, "Сообщите приглашающему код:\n"+
						strconv.Itoa(int(message.Message.Chat.ID))))
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
				if message.Message.Chat.Type != supergroupChatType {
					// TODO: пишем в лог
					continue
				}

				// Цикл проверок
				for _, v := range sc.policies {
					if v.Check(message) != nil {
						// Если проверка сработала, то удаляем сообщение
						dm := tgbotapi.NewDeleteMessage(message.Message.Chat.ID, message.Message.MessageID)
						if _, fail := botApi.Request(dm); fail != nil {
							// TODO: сохраняем кучу данных в лог
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
