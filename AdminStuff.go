package main

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
	"moderatorBot/internal/policy"
	"moderatorBot/internal/storage"
	"strconv"
	"strings"
	"time"
)

const (
	ReqChatIdText       = "Введите токен нового админа\nЕсли у вас нет токена, попросите приглашённого написать боту любой текст"
	ReqNameText         = "Введите имя нового админа"
	ReqConfirmationText = "Проверьте токен и имя нового админа"
	ReqBanWordText      = "Введите запрещённое слово"
	ReqWarningsValText  = "Введите количество предупреждений перед мутом пользователя"
	ReqBanTimeText      = "Введите врмя мута в минутах"

	ConfirmText    = "Подтвердить"
	NotConfirmText = "Повторный ввод"
	BanWordAdded   = "Запрещённое слово успешно добавленно"

	NotAdminText = "Вы не являетесь админом\nЕсли вы хотите, чтобы вас назначили админом, сообщите приглашающему код:\n"

	AddBannedWordText    = "Добавить запрещённое слово"
	AddAdminText         = "Добавить админа"
	SetBanTimeText       = "Установить время мута"
	SetWarningsText      = "Установить количество предупреждений"
	GetSettingsText      = "Узнать настройки"
	DeleteBannedWordText = "Удалить запрещённое слово"

	WhatToDoText = "Что делать будем?"
)

var MainAdminKeyboard = tgbotapi.NewReplyKeyboard(
	tgbotapi.NewKeyboardButtonRow(tgbotapi.NewKeyboardButton(AddBannedWordText), tgbotapi.NewKeyboardButton(SetBanTimeText)),
	tgbotapi.NewKeyboardButtonRow(tgbotapi.NewKeyboardButton(DeleteBannedWordText), tgbotapi.NewKeyboardButton(SetWarningsText)),
	tgbotapi.NewKeyboardButtonRow(tgbotapi.NewKeyboardButton(AddAdminText), tgbotapi.NewKeyboardButton(GetSettingsText)),
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
		hideKeyboard(id, ReqBanWordText)
		tg, name := CreateAdmin(id, channel)
		storage.AddAdmins(tg, name)
		BotAPI.Send(tgbotapi.NewMessage(id, "Админ с токеном: "+strconv.Itoa(int(tg))+"\nИменем: "+name+"\nДобавлен!"))
		showKeyboard(id, WhatToDoText, MainAdminKeyboard)
	} else {
		BotAPI.Send(tgbotapi.NewMessage(id, NotAdminText+strconv.Itoa(int(id))))
	}
}

func BanWordAddition(id int64, channel chan tgbotapi.Update, storage storage.Interface) {
	if IsItAdmin(id, storage) == true {
		hideKeyboard(id, ReqBanWordText)
		banWord, err := InputText(id, channel, "")
		if err != nil {
			//TODO: какая то реакция
		}
		storage.AddBannedWord(strings.ToLower(banWord))
		log.Println(len(ContainsPolicy))
		ContainsPolicy = append(ContainsPolicy, policy.NewContains(banWord))
		BotAPI.Send(tgbotapi.NewMessage(id, BanWordAdded))
		showKeyboard(id, WhatToDoText, MainAdminKeyboard)
	} else {
		BotAPI.Send(tgbotapi.NewMessage(id, NotAdminText+strconv.Itoa(int(id))))
	}
}

func SetWarningsVal(id int64, storage storage.Interface, channel chan tgbotapi.Update) {
	if IsItAdmin(id, storage) == true {
		hideKeyboard(id, ReqWarningsValText)
		panishments.Warnings = trollcheck(id, channel)
		storage.SetWarnings(panishments.Warnings)
		showKeyboard(id, WhatToDoText, MainAdminKeyboard)
	}
}

func setBantime(id int64, storage storage.Interface, channel chan tgbotapi.Update) {
	if IsItAdmin(id, storage) == true {
		hideKeyboard(id, ReqBanTimeText)
		panishments.Bandur = time.Minute * time.Duration(trollcheck(id, channel))
		storage.SetBanTime(trollcheck(id, channel))
		showKeyboard(id, WhatToDoText, MainAdminKeyboard)
	}
}

func trollcheck(id int64, channel chan tgbotapi.Update) int {
	for {
		val, err := InputText(id, channel, "")
		if err != nil {
			//TODO:я обязательно сделаю все todo...
		}
		if strconv.Atoi(val); err != nil {
			BotAPI.Send(tgbotapi.NewMessage(id, "ЧИСЛО! ЧИСЛО Я СКАЗАЛ!"))
			BotAPI.Send(tgbotapi.NewMessage(id, "ЛЮДИ, ВЫ ВООБЩЕ УМЕЕТЕ ОТЛИЧАТЬ ЧИСЛО ОТ ОСТАЛЬНЫХ СИМВОЛОВ?!"))
			BotAPI.Send(tgbotapi.NewMessage(id, "Ну реально, не тупите, люди, прошу число, значит число!"))
			BotAPI.Send(tgbotapi.NewMessage(id, "Если бы я не проверил, ты бы всё сломал"))
			BotAPI.Send(tgbotapi.NewMessage(id, "ВСЁ СЛОМАЛ!"))
			BotAPI.Send(tgbotapi.NewMessage(id, "Я БЫ УМЕР!"))
			BotAPI.Send(tgbotapi.NewMessage(id, "Ладно, я бы не умер, надо было бы всего лишь перезапустить..."))
			BotAPI.Send(tgbotapi.NewMessage(id, "Но всё равно! ОНО ТЕБЕ НАДО!"))
			BotAPI.Send(tgbotapi.NewMessage(id, "Поэтому введи пожалуйста теперь число и не беси меня"))
		} else {
			intval, _ := strconv.Atoi(val)
			if intval <= 0 {
				BotAPI.Send(tgbotapi.NewMessage(id, "БОЖЕ, СЕРЬЁЗНО?!"))
				BotAPI.Send(tgbotapi.NewMessage(id, "ЭТО ТЫ ТИПА РЕШИЛ ТАК ПОТЕСТИТЬ?"))
				BotAPI.Send(tgbotapi.NewMessage(id, "Ооо... А что же будет, если я введу отрицательное число? Хм..."))
				BotAPI.Send(tgbotapi.NewMessage(id, "А НИЧЕГО НЕ БУДЕТ! НИ-ЧЕ-ГО!"))
				BotAPI.Send(tgbotapi.NewMessage(id, "Возможно, не отрицательное, а ноль, проверка одна и та же, я хз"))
				BotAPI.Send(tgbotapi.NewMessage(id, "А, ты ввёл "+val+". Как я и сказал, это не подходит"))
				BotAPI.Send(tgbotapi.NewMessage(id, "Хватит меня ломать! Я и так на добром слове держусь"))
			} else {
				return intval
			}
		}
	}
}

func GetPanishments(id int64, storage storage.Interface) {
	if IsItAdmin(id, storage) == true {
		war, ban := storage.GetPanishments()
		BotAPI.Send(tgbotapi.NewMessage(id, "Время мута (минуты): "+ban+"\nКоличество предупреждений: "+war))
	}
}

func DeleteBannedWord(id int64, channel chan tgbotapi.Update, storage storage.Interface) []policy.Interface {
	if IsItAdmin(id, storage) == true {
		word, err := InputText(id, channel, "Введите слово для удаления")
		if err != nil {
			log.Println(err.Error())
		}
		var r []policy.Interface
		for i := 0; i < len(ContainsPolicy); i++ {
			if ContainsPolicy[i].GetContains() == word {
				log.Println(ContainsPolicy[0:i], ContainsPolicy[i+1:], "+++", i)
				r = append(ContainsPolicy[0:i], ContainsPolicy[i:]...)
				break
			}
		}

		if storage.DeleteBannedWord(word) == true {
			BotAPI.Send(tgbotapi.NewMessage(id, "'"+word+"' удалено"))
		} else {
			BotAPI.Send(tgbotapi.NewMessage(id, "Слово '"+word+"' не найдено"))
		}
		return r
	}
	return ContainsPolicy
}

//func RefreshPolicy(storage storage.Interface) {
//	myPolicy, err := storage.GetPolicy()
//	if err != nil {
//		log.Println(err.Error())
//	}
//	ContainsPolicy = ContainsPolicy[:0]
//	for i := 0; i < len(myPolicy); i++ {
//		ContainsPolicy = append(ContainsPolicy, policy.NewContains(myPolicy[i]))
//	}
//}
