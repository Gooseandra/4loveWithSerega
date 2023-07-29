package policy

import (
	"errors"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
	"strings"
)

type Urls struct {
	text    string
	exclude []Contains
}

func NewOkUrl(text string) Urls { return Urls{text: text} }

func (c Urls) Check(update tgbotapi.Update) error {
	if update.Message != nil {
		words := strings.Split(strings.ToLower(update.Message.Text), " ")
		for index := 0; index < len(words); index++ {
			word := words[index]
			if len(word) > 4 {
				if word[:4] == "www." {
					slashes := strings.Split(update.Message.Text, "/")
					for i := 0; i < len(slashes); i++ {
						if slashes[0] == "www."+c.text {
							break
						}
						if i == len(slashes)-1 {
							return errors.New("Содержит ссылку")
						}
					}
				} else if word[:5] == "https" {
					slashes := strings.Split(update.Message.Text, "/")
					for i := 0; i < len(slashes); i++ {
						log.Println("https://" + slashes[i])
						log.Println(c.text)
						if slashes[i] == c.text || slashes[i] == "https://www."+c.text {
							break
						}
						if i == len(slashes)-1 {
							return errors.New("Содержит ссылку")
						}
					}
				}
			}
		} // word[:5] == "https"
	}
	return nil
}

func (c Urls) GetContains() string {
	return c.text
}
