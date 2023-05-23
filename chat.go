package main

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
	"moderatorBot/internal/policy"
	"moderatorBot/internal/storage"
	"strconv"
	"sync"
	"time"
)

const (
	MainStatus     = ChatStatus(iota)
	TrainStatus    = ChatStatus(iota)
	SettingsStatus = ChatStatus(iota)
	nilStatus      = ChatStatus(iota)
	showStatus     = ChatStatus(iota)
)

type (
	BaseChat struct {
		channel chan tgbotapi.Update
		db      storage.ChatIdModel
		tg      int64
		//timeStart  time.Time
		//timeFinish time.Time
	}

	ChatInterface interface {
		routine(*tgbotapi.BotAPI, map[int64]ChatInterface, *sync.Mutex, storage.Interface)
		send(tgbotapi.Update)
	}

	ChatStatus byte

	PrivateChat struct{ BaseChat }

	SupergroupChat struct {
		BaseChat
		policies  []policy.Interface
		moderated bool
	}
)

func (bc BaseChat) send(update tgbotapi.Update) { bc.channel <- update }

func (pc PrivateChat) routine(_ *tgbotapi.BotAPI, chats map[int64]ChatInterface, mainMutex *sync.Mutex, storage storage.Interface) {
	lastMassageTime := time.After(time.Hour * 10)
	status := nilStatus
	for {
		select {
		case message := <-pc.channel:
			//close(lastMassageTime)
			lastMassageTime = time.After(time.Hour * 10)
			log.Println("MainStatus.Update")
			if message.Message != nil {
				fmt.Println(message.Message.Chat.Type)
				switch status {
				case nilStatus:
					switch message.Message.Text {
					case "/start":

					}
				case MainStatus:

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
				if message.Message.Chat == nil {
					// TODO: пишем в лог
					continue
				}
				if message.Message.Chat.Type != supergroupChatType {
					// TODO: пишем в лог
					continue
				}
				for _, v := range sc.policies {
					if v.Check(message) != nil {
						dm := tgbotapi.NewDeleteMessage(message.Message.Chat.ID, message.Message.MessageID)
						if _, fail := botApi.Request(dm); fail != nil {
							// TODO: сохраняем кучу данных в лог
						}
						break
					}
				}
			default:
				// TODO: пишем message в лог
			}
		case <-lastMassageTime:
			log.Println("time out")
			mainMutex.Lock()
			log.Println("Chat " + strconv.FormatInt(sc.tg, 10) + " deleted")
			log.Println(chats)
			delete(chats, sc.tg)
			mainMutex.Unlock()
		}
	}
}
