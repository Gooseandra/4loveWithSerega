package main

import (
	"errors"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
	"strconv"
)

const (
	ReqChatIdText       = "Введите токен нового админа\nЕсли у вас нет токена, попросите приглашённого написать боту команду add me"
	ReqNameText         = "Введите имя нового админа"
	ReqConfirmationText = "Проверьте токен и имя нового админа"

	ConfirmText    = "Подтвердить"
	NotConfirmText = "Повторный ввод"
)

var ConfirmationKeyboard = tgbotapi.NewReplyKeyboard(
	tgbotapi.NewKeyboardButtonRow(tgbotapi.NewKeyboardButton(ConfirmText)),
	tgbotapi.NewKeyboardButtonRow(tgbotapi.NewKeyboardButton(NotConfirmText)),
)

func InputText(id int64, channel chan tgbotapi.Update, disc string) (string, error) {
	BotAPI.Send(tgbotapi.NewMessage(id, disc))
	message := <-channel
	if message.Message == nil {
		return "", errors.New("Не то ввел, братишка")
	}
	return message.Message.Text, nil
}

func CreateAdmin(id int64, channel chan tgbotapi.Update) (int64, string) {
	log.Println(id)
	for {
		tg, err := InputText(id, channel, ReqChatIdText)
		if err != nil {
			//TODO: пишем в лог
		}
		name, err := InputText(id, channel, ReqNameText)
		if err != nil {
			//TODO: пишем в лог
		}
		showCmd := tgbotapi.NewMessage(id, ReqConfirmationText+"\nТокен: "+tg+"\nИмя: "+name)
		showCmd.ReplyMarkup = ConfirmationKeyboard
		BotAPI.Request(showCmd)
		confirmation, err := InputText(id, channel, "")
		if confirmation == ConfirmText {
			temp, err := strconv.Atoi(tg)
			if err != nil {
				log.Println(err.Error())
			}
			return int64(temp), name
		}
	}
}
