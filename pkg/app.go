package bot

import (
	"fmt"
	"github.com/line/line-bot-sdk-go/linebot"
	"log"
	"net/http"
	"regexp"
	"strings"
)

type LineDictionaryApp struct {
	bot *linebot.Client
}

func NewLineDictionaryApp(lineConfiguration LineConfiguration) (*LineDictionaryApp, error) {
	bot, err := linebot.New(lineConfiguration.LineSecret, lineConfiguration.LineToken)
	if err != nil {
		return nil, err
	}

	return &LineDictionaryApp{
		bot: bot,
	}, nil
}

func (app LineDictionaryApp) WebHook() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		events, err := app.bot.ParseRequest(r)
		if err != nil {
			if err == linebot.ErrInvalidSignature {
				w.WriteHeader(400)
			} else {
				w.WriteHeader(500)
			}
			return
		}
		for _, event := range events {
			log.Printf("%+v", event)
			if event.Type == linebot.EventTypeMessage {
				switch message := event.Message.(type) {
				case *linebot.TextMessage:
					if err := app.processText(message, event.ReplyToken); err != nil {
						log.Println(err)
					}
				}
			}
		}
	}
}

func (app LineDictionaryApp) processText(message *linebot.TextMessage, replyToken string) error {
	fmt.Println(message.Text)
	return nil
}

func detectFirstWord(message string) string {
	token := strings.Split(strings.TrimSpace(message), " ")
	if ok, _ := regexp.MatchString("[a-zA-Z]+", token[0]); ok {
		return token[0]
	}
	return ""
}