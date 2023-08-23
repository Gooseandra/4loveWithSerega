package main

import (
	"encoding/json"
	"fmt"
	"github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/go-yaml/yaml"
	_ "github.com/lib/pq"
	"log"
	"moderatorBot/internal/policy"
	"moderatorBot/internal/storage"
	"net/http"
	"os"
	"strconv"
	"sync"
	"time"
)

const (
	privateChatType    = "private"
	supergroupChatType = "supergroup"
	groutChatType      = "group"

	addres = "127.0.0.1:80"
)

var BotAPI *tgbotapi.BotAPI

var ContainsPolicy []policy.Interface
var UrlPolicy []policy.Interface
var whiteList []string

func startServer(storage storage.Interface) {
	log.Println("дада ")
	http.HandleFunc("/", func(rw http.ResponseWriter, req *http.Request) {
		admins, _ := storage.LoadAdmins()

		f := "3"
		ttt := "hjhjh"
		ttt = ttt + "yyjy" + f
		log.Println(admins)
		//fmt.Fprintf(rw, "<html><body><div>%s, %s</div></body></html>", admins[0].Tg, admins[1].Tg)
		fmt.Fprintf(rw, "<html>"+
			"<body>"+
			"<div><input id='input'><Br>"+
			"<button onclick=\"let k = document.getElementById('input').value;fetch('/test?id=' + k).then(async function (res) {let tmp = await res.text(); alert(tmp)}); \">Казнить</button><br>"+
			"<button onclick=\"let k = document.getElementById('input').value;fetch('/test2?id=' + k).then(async function (res) {let tmp = await res.text(); alert(tmp)}); \">Помиловать</button><br>"+
			"<button onclick=\"let k = document.getElementById('input').value;fetch('/bannedusers?id=' + k).then(async function (res) {let tmp = await res.text(); console.log(tmp)}); \">Кто ЗДОХЪ?</button>"+
			"</div>"+
			"<div id='bann'>"+
			"</div>"+
			"<script>"+
			"fetch('/bannedusers').then(async function (res){"+
			"let tmp = await res.json();"+
			"let el = document.getElementById('bann');"+
			"for (let item of tmp){"+
			"const button = document.createElement(\"button\");"+
			"button.innerHTML = 'Помиловать ' + item;"+
			"button.addEventListener('click',()=>{"+
			"console.log(item);"+
			"fetch('/test2?id=' + item).then(async function (res) {let tmp = await res.text(); alert(tmp)});"+
			"});"+
			"el.appendChild(button);"+
			"}"+
			"console.log(tmp);"+
			"})"+
			"</script>"+
			"</body>"+
			"</html>")
	})
	http.HandleFunc("/test", func(rw http.ResponseWriter, req *http.Request) {
		id := req.FormValue("id")
		log.Println(id)
		//admins, _ := storage.LoadAdmins()
		//log.Println(admins)
		//fmt.Fprintf(rw, "<html><body><div>uuuuu=%s, %s</div></body></html>", admins[0].Tg, admins[1].Tg)
		banned := storage.GetBanList()
		var found = false
		for _, v := range banned {
			if v == id {
				found = true
				//fmt.Fprintf(rw, "")
				break
			}
		}
		if found == false {
			intid, _ := strconv.Atoi(id)
			storage.Crime(int64(intid), 0, time.Minute*30)
			storage.Crime(int64(intid), 0, time.Minute*30)
			fmt.Fprintf(rw, "ok")
		} else {
			fmt.Fprintf(rw, "err")
		}
	})
	http.HandleFunc("/test2", func(rw http.ResponseWriter, req *http.Request) {
		//admins, _ := storage.LoadAdmins()
		id := req.FormValue("id")
		log.Println(id)
		//log.Println(admins)
		//fmt.Fprintf(rw, "<html><body><div>uuuuu=%s, %s</div></body></html>", admins[0].Tg, admins[1].Tg)
		banned := storage.GetBanList()
		var found = false
		for _, v := range banned {
			if v == id {
				found = true
				//fmt.Fprintf(rw, "")
				break
			}
		}
		if found == true {
			intid, _ := strconv.Atoi(id)
			storage.Unban(int64(intid), time.Nanosecond)
			fmt.Fprintf(rw, "ok")
		} else {
			fmt.Fprintf(rw, "err")
		}
	})
	http.HandleFunc("/bannedusers", func(rw http.ResponseWriter, req *http.Request) {
		banned := storage.GetBanList()
		res, _ := json.Marshal(banned)
		fmt.Fprintf(rw, "%s", res)
	})
	log.Fatal("HTTP server error: ", http.ListenAndServe(addres, nil))
}

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
	go startServer(storage_)

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
				if message.Message.Chat.Type == privateChatType {
					var model storage.UpsertUserByTgModel
					model, fail = storage_.UpsertUserByTg(message.Message.Chat.ID, message.Message.Chat.Title)
					chat = PrivateChat{BaseChat: BaseChat{
						channel: make(chan tgbotapi.Update),
						db:      model.Id,
						tg:      message.FromChat().ID}}
				} else if message.Message.Chat.Type == supergroupChatType || message.Message.Chat.Type == groutChatType {
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
						UrlPolicy = append(UrlPolicy, policy.NewOkUrl(myUrls[i]))
					}
					whiteList = storage_.GetWhiteList()
					//policies = ContainsPolicy
					chat = SupergroupChat{
						BaseChat: BaseChat{
							channel: make(chan tgbotapi.Update),
							db:      model.Id,
							tg:      message.FromChat().ID},
						moderated: model.Moderated}
				} else {
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
