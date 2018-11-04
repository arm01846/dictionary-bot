package bot

import (
	"github.com/line/line-bot-sdk-go/linebot"
	"log"
	"net/http"
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
		}
	}
}