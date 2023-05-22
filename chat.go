package main

import (
	"fmt"
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
	BaseChat struct {
		id         int64
		channel    chan tgbotapi.Update
		timeStart  time.Time
		timeFinish time.Time
	}

	ChatInterface interface {
		routine(map[int64]ChatInterface, *sync.Mutex, storage.Interface)
		send(tgbotapi.Update)
	}

	ChatStatus byte

	PrivateChat struct{ BaseChat }

	SupergroupChat struct{ BaseChat }
)

func (bc BaseChat) send(update tgbotapi.Update) { bc.channel <- update }

func (pc PrivateChat) routine(chats map[int64]ChatInterface, mainMutex *sync.Mutex, storage storage.Interface) {
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
			log.Println("Chat " + strconv.FormatInt(pc.id, 10) + " deleted")
			log.Println(chats)
			delete(chats, pc.id)
			mainMutex.Unlock()
		}
	}
}

func (sc SupergroupChat) routine(chats map[int64]ChatInterface, mainMutex *sync.Mutex, storage storage.Interface) {
	lastMassageTime := time.After(time.Hour * 10)
	status := nilStatus
	for {
		select {
		case message := <-sc.channel:
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
			log.Println("Chat " + strconv.FormatInt(sc.id, 10) + " deleted")
			log.Println(chats)
			delete(chats, sc.id)
			mainMutex.Unlock()
		}
	}
}
