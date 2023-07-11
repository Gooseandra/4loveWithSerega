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
	supergroupChatType = "group"
)

var BotAPI *tgbotapi.BotAPI

var ContainsPolicy []policy.Interface
var UrlPolicy []policy.Interface

func main() {
	var mainMutex sync.Mutex
	var chats = map[int64]ChatInterface{}
	var settings Settings

	var storage_ storage.Interface
	bytes, fail := os.ReadFile(".yml")
	if fail != nil {
		log.Println(fail.Error())
		log.Panic(fail.Error())
	}
	fail = yaml.Unmarshal([]byte(bytes), &settings)
	if fail != nil {
		log.Panic(fail.Error())
	}
	log.Println(settings)
	storage_, fail = storage.NewPostgres(settings.Database.Arguments)
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

	panishments.Bandur = storage_.GetBanTime()
	panishments.Warnings = storage_.GetWarnings()

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
			log.Println(message.Message.Chat.Type)
			mainMutex.Lock()
			chat, found := chats[message.Message.Chat.ID]
			if !found {
				switch message.Message.Chat.Type {
				case privateChatType:
					var model storage.UpsertUserByTgModel
					model, fail = storage_.UpsertUserByTg(message.Message.Chat.ID, message.Message.Chat.Title)
					chat = PrivateChat{BaseChat: BaseChat{
						channel: make(chan tgbotapi.Update),
						db:      model.Id,
						tg:      message.FromChat().ID}}
				case supergroupChatType:
					log.Println("okok")
					var model storage.UpsertChatByTgModel
					model, fail = storage_.UpsertChatByTg(message.Message.Chat.ID, message.Message.Chat.Title)
					if fail != nil {
						// TODO: пишем в лог, возможно обрабатываем ощибку недоступности БД
						//continue
					}
					myPolicy, err := storage_.GetPolicy()
					myUrls := storage_.GetUrls()
					if err != nil {
						//TODO: пишем в лог
					}
					for i := 0; i < len(myPolicy); i++ {
						ContainsPolicy = append(ContainsPolicy, policy.NewContains(myPolicy[i]))
					}
					for i := 0; i < len(myUrls); i++ {
						ContainsPolicy = append(ContainsPolicy, policy.NewOkUrl(myUrls[i]))
					}
					chat = SupergroupChat{
						BaseChat: BaseChat{
							channel: make(chan tgbotapi.Update),
							db:      model.Id,
							tg:      message.FromChat().ID},
						moderated: model.Moderated,
						policies:  ContainsPolicy}
				default:
					// TODO: пишем в лог
					continue
				}
				chats[message.FromChat().ID] = chat
				log.Println("Chat " + strconv.FormatInt(message.FromChat().ID, 10) + " created")
				go chat.routine(BotAPI, chats, &mainMutex, storage_)
			}
			mainMutex.Unlock()
			chat.send(message)
		}
	}
}
