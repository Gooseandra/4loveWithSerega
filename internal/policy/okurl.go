package policy

import (
	"errors"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"strings"
)

type StartWith struct {
	text string
}

func okUrl(text string) Contains { return Contains{text: text} }

func (sw StartWith) Check(update tgbotapi.Update) error {
	if update.Message != nil {
		if strings.HasPrefix(update.Message.Text, sw.text) {
			return errors.New("Содержит словосочетание '" + sw.text + " в начале'")
		}
	}
	return nil
}
