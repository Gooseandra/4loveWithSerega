package main

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
	"moderatorBot/internal/policy"
	"moderatorBot/internal/storage"
	"strconv"
)

const (
	ReqChatIdText       = "Введите токен нового админа\nЕсли у вас нет токена, попросите приглашённого написать боту команду add me"
	ReqNameText         = "Введите имя нового админа"
	ReqConfirmationText = "Проверьте токен и имя нового админа"
	ReqBanWordText      = "Введите запрещённое слово"

	ConfirmText    = "Подтвердить"
	NotConfirmText = "Повторный ввод"
	BanWordAdded   = "Запрещённое слово успешно добавленно"

	NotAdminText = "Вы не являетесь админом\nЕсли вы хотите, чтобы вас назначили админом, сообщите приглашающему код:\n"

	AddBannedWordText = "Добавить запрещённое слово"
	AddAdminText      = "Добавить админа"

	WhatToDoText = "Что делать будем?"
)

var MainAdminKeyboard = tgbotapi.NewReplyKeyboard(
	tgbotapi.NewKeyboardButtonRow(tgbotapi.NewKeyboardButton(AddBannedWordText)),
	tgbotapi.NewKeyboardButtonRow(tgbotapi.NewKeyboardButton(AddAdminText)),
)

var ConfirmationKeyboard = tgbotapi.NewReplyKeyboard(
	tgbotapi.NewKeyboardButtonRow(tgbotapi.NewKeyboardButton(ConfirmText)),
	tgbotapi.NewKeyboardButtonRow(tgbotapi.NewKeyboardButton(NotConfirmText)),
)

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
		showKeyboard(id, ReqConfirmationText+"\nТокен: "+tg+"\nИмя: "+name, ConfirmationKeyboard)
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

func AdminAddition(id int64, channel chan tgbotapi.Update, storage storage.Interface) {
	if IsItAdmin(id, storage) == true {
		tg, name := CreateAdmin(id, channel)
		storage.AddAdmins(tg, name)
	} else {
		BotAPI.Send(tgbotapi.NewMessage(id, NotAdminText+strconv.Itoa(int(id))))
	}
}

func BanWordAddition(id int64, channel chan tgbotapi.Update, storage storage.Interface, ContainsPolicy []policy.Interface) {
	if IsItAdmin(id, storage) == true {
		hideKeyboard(id, ReqBanWordText)
		banWord, err := InputText(id, channel, "")
		if err != nil {
			//TODO: какая то реакция
		}
		storage.AddBannedWord(banWord)
		ContainsPolicy = append(ContainsPolicy, policy.NewContains(banWord))
		BotAPI.Send(tgbotapi.NewMessage(id, BanWordAdded))
		showKeyboard(id, WhatToDoText, MainAdminKeyboard)
	} else {
		BotAPI.Send(tgbotapi.NewMessage(id, NotAdminText+strconv.Itoa(int(id))))
	}
}
