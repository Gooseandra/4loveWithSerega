package main

import (
	"github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/go-yaml/yaml"
	_ "github.com/lib/pq"
	"log"
	"moderatorBot/internal/storage"
	"os"
	"strconv"
	"sync"
)

var BotAPI *tgbotapi.BotAPI

func main() {
	var mainMutex sync.Mutex
	var chats = map[int64]Chat{}
	var settings Settings
	bytes, fail := os.ReadFile(".yml")
	if fail != nil {
		log.Panic(fail.Error())
	}
	fail = yaml.Unmarshal([]byte(bytes), &settings)
	if fail != nil {
		log.Panic(fail.Error())
	}
	log.Println(settings)
	s, fail := storage.NewPostgres(settings.Database.Arguments)
	if fail != nil {
		log.Panic(fail.Error())
	}
	BotAPI, fail = tgbotapi.NewBotAPI(settings.Telegram)
	if fail != nil {
		log.Panic(fail)
	}
	update := tgbotapi.NewUpdate(0)
	update.Timeout = settingsTimeout
	channel := BotAPI.GetUpdatesChan(update)
	//admins, fail := NewAdmins(s)
	//if fail != nil {
	//	log.Panic(fail)
	//}
	for {
		select {
		case message := <-channel:
			mainMutex.Lock()
			chat, found := chats[message.FromChat().ID]
			if !found {
				chat = Chat{id: message.FromChat().ID, channel: make(chan tgbotapi.Update)}
				chats[message.FromChat().ID] = chat
				log.Println("Chat " + strconv.FormatInt(message.FromChat().ID, 10) + " created")
				go chat.routine(chats, &mainMutex, s)
			}
			mainMutex.Unlock()
			chat.channel <- message
		}
	}
}
