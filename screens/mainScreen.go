package screens

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

const (
	EnrollText          = "Записаться на сеанс"
	GetRecordsText      = "Просмотреть существующие записи"
	GetContactsText     = "Контакты/обратная связь с (кем?)"
	MainScreenHelloText = "s"
)

var mainScreenKB = tgbotapi.NewReplyKeyboard(
	tgbotapi.NewKeyboardButtonRow(tgbotapi.NewKeyboardButton("EnrollText")),
	tgbotapi.NewKeyboardButtonRow(tgbotapi.NewKeyboardButton("GetRecordsText")),
	tgbotapi.NewKeyboardButtonRow(tgbotapi.NewKeyboardButton("GetContactsText")),
)

const (
	MainStatus = ChatStatus(iota)
)

var BotAPI *tgbotapi.BotAPI

type ChatStatus uint8

func ShowMainScreen(id int64) ChatStatus {
	showCmd := tgbotapi.NewMessage(id, "MainScreenHelloText")
	showCmd.ReplyMarkup = mainScreenKB
	BotAPI.Request(showCmd)
	return MainStatus
}

// тут и в enrollScreen представленны два вида, как мы можем перейти на экран
// либо создаём всё тут, активируем кнопки и возвращаем статус, либо, как сделанно в enrollScreen
// создаём конфиг сообщения и возвращаем его, а отсылаем (Request()) делаем в рутине
// скажи, что лучше, что больше нравится, что практичнее и правильнее
// я сколняюсь в сторону возврата статуса
