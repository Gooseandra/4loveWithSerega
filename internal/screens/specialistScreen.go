package screens

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

//Пользователь выбрал желаемое заведение:
//Выдается меню из 3 кнопок:
//1. "Я знаю - к кому хочу!"
//2. "Мне все равно, запишите меня к кому угодно"

// НАЙС 3 КНОПКИ :)))

const (
	IKnowWhoAmIWantText = "Я знаю - к кому хочу!"
	NevermindText       = "Мне все равно, запишите меня к кому угодно"
	SpecialistHelloText = "Хотите записаться к конкретному мастеру?"
	SpecialistStatus    = ChatStatus(iota)
)

var doUKnowWhoUWantKB = tgbotapi.NewReplyKeyboard(
	tgbotapi.NewKeyboardButtonRow(tgbotapi.NewKeyboardButton("IKnowWhoAmIWantText")),
	tgbotapi.NewKeyboardButtonRow(tgbotapi.NewKeyboardButton("NevermindText")),
)

func doUKnowWhoUWant(id int64) ChatStatus {
	showCmd := tgbotapi.NewMessage(id, "SpecialistHelloText")
	showCmd.ReplyMarkup = mainScreenKB
	BotAPI.Request(showCmd)
	return SpecialistStatus
}

// дальше я решил пойти спать, хочется отметить, что в тз написано, что когда челик выбирает,
// что он знает, к кому хочет, сначала выбирвается дата, а потом специалист, который в этот день
// работает, я так понимаю, это опечатка и сначала выбирается специалист, а потом пусть высвечивается
// дата, когда у него есть запись

// upd: почитал тз дальше, там идёт уже правильно (в пункте 3), что сначала выбирает сотрудника, а
// потом время, кароч, я просто доебаться решил :)
