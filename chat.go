package main

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
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
	Chat struct {
		id         int64
		channel    chan tgbotapi.Update
		timeStart  time.Time
		timeFinish time.Time
	}

	ChatStatus byte
)

func (chat Chat) routine(chats map[int64]Chat, mainMutex *sync.Mutex, storage storage.Interface) {
	lastMassageTime := time.After(time.Hour * 10)
	status := nilStatus
	for {
		select {
		case message := <-chat.channel:
			//close(lastMassageTime)
			lastMassageTime = time.After(time.Hour * 10)
			log.Println("MainStatus.Update")
			if message.Message != nil {
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
			log.Println("Chat " + strconv.FormatInt(chat.id, 10) + " deleted")
			log.Println(chats)
			delete(chats, chat.id)
			mainMutex.Unlock()
		}
	}
}
