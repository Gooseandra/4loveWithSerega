package main

import (
	"github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/go-yaml/yaml"
	_ "github.com/lib/pq"
	"log"
	"moderatorBot/internal/policy"
	"moderatorBot/internal/storage"
	"os"
	"strconv"
	"sync"
)

const (
	privateChatType    = "private"
	supergroupChatType = "supergroup"
)

var BotAPI *tgbotapi.BotAPI

func main() {
	var mainMutex sync.Mutex
	var chats = map[int64]ChatInterface{}
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
			if message.Message == nil {
				// TODO: пишем в лог
				continue
			}
			if message.Message.Chat == nil {
				// TODO: пишем в лог
				continue
			}
			mainMutex.Lock()
			chat, found := chats[message.Message.Chat.ID]
			if !found {
				var id storage.ChatIdModel
				if id, fail = s.UpdateChatByTg(message.Message.Chat.ID, message.Message.Chat.Title); fail != nil {
					// TODO: пишем в лог, возможно обрабатываем ощибку недоступности БД
					continue
				}
				baseChat := BaseChat{channel: make(chan tgbotapi.Update), db: id, tg: message.FromChat().ID}
				switch message.Message.Chat.Type {
				case privateChatType:
					chat = PrivateChat{BaseChat: baseChat}
				case supergroupChatType:
					chat = SupergroupChat{
						BaseChat: baseChat,
						policies: []policy.Interface{policy.NewContains("asd")}}
				default:
					// TODO: пишем в лог
					continue
				}
				chats[message.FromChat().ID] = chat
				log.Println("Chat " + strconv.FormatInt(message.FromChat().ID, 10) + " created")
				go chat.routine(BotAPI, chats, &mainMutex, s)
			}
			mainMutex.Unlock()
			chat.send(message)
		}
	}
}
