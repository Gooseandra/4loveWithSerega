package policy

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

type Interface interface {
	Check(tgbotapi.Update) error
}
