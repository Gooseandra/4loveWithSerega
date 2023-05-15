package screens

import (
	"database/sql"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
)

//Меню "Записаться на сеанс":
//1. Пользователю выдается список заведений, в которых оказывается услуга.
//2. Вернуться в главное меню.

const (
	choseInstitutionText = "Выберите заведение"
	BackText             = "Назад"
	EnrollStatus         = ChatStatus(iota)
)

func findInstitutions(id int64, db *sql.DB) tgbotapi.MessageConfig {
	showCmd := tgbotapi.NewMessage(id, choseInstitutionText)
	Keyboards := [][]tgbotapi.KeyboardButton{}
	rows, err := db.Query(`select * from test`) // вот тут меняем, для себя просто пометил, потому что потом потеряю
	if err != nil {
		log.Println(err.Error())
	}
	var str string
	for rows.Next() {
		rows.Scan(&str)
		log.Println(str)
		Keyboards = append(Keyboards, []tgbotapi.KeyboardButton{tgbotapi.NewKeyboardButton(str)})
	}
	Keyboards = append(Keyboards, []tgbotapi.KeyboardButton{tgbotapi.NewKeyboardButton(BackText)})
	showCmd.ReplyMarkup = tgbotapi.NewReplyKeyboard(Keyboards...)
	return showCmd
}
