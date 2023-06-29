package main

import (
	"errors"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"moderatorBot/internal/storage"
	"time"
)

var panishments struct {
	Warnings int
	Bandur   time.Duration
}

func showKeyboard(id int64, disc string, kb tgbotapi.ReplyKeyboardMarkup) {
	showCmd := tgbotapi.NewMessage(id, disc)
	showCmd.ReplyMarkup = kb
	BotAPI.Request(showCmd)
}

func hideKeyboard(id int64, disc string) {
	msg := tgbotapi.NewMessage(id, disc)
	msg.ReplyMarkup = tgbotapi.NewRemoveKeyboard(true)
	BotAPI.Request(msg)
}

func InputText(id int64, channel chan tgbotapi.Update, disc string) (string, error) {
	BotAPI.Send(tgbotapi.NewMessage(id, disc))
	message := <-channel
	if message.Message == nil {
		return "", errors.New("Не то ввел, братишка")
	}
	return message.Message.Text, nil
}

func IsItAdmin(id int64, storage storage.Interface) bool {
	admins, err := storage.LoadAdmins()
	if err != nil {
		//TODO: пришем в лог
	}
	for i := 0; i < len(admins); i++ {
		if admins[i].Tg == id {
			return true
		}
	}
	return false
}
