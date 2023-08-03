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
		var wordsAndStrings []string
		for i := 0; i < len(words); i++ {
			temp := strings.Split(words[i], "\n")
			for _, v := range temp {
				wordsAndStrings = append(wordsAndStrings, v)
			}
		}
		for index := 0; index < len(wordsAndStrings); index++ {
			word := wordsAndStrings[index]
			if strings.ToLower(word) == c.text {
				return errors.New("Содержит слово '" + c.text + "'")
			}
		} // word[:5] == "https"
	}
	return nil
}

func (c Contains) GetContains() string {
	return c.text
}
