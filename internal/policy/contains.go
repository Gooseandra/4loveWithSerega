package policy

import (
	"errors"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"strings"
)

type Contains struct {
	text    string
	exclude []Contains
}

func NewContains(text string) Contains { return Contains{text: text} }

func (c Contains) Check(update tgbotapi.Update) error {
	if update.Message != nil {
		words := strings.Split(update.Message.Text, " ")
		for index := 0; index < len(words); index++ {
			if words[index] == c.text {
				return errors.New("Содержит словосочетание '" + c.text + "'")
			}
		}
	}
	return nil
}
