package main

import (
	"github.com/arm01846/dictionary-bot/pkg"
	"log"
	"net/http"
	"os"
)

func main() {
	port := os.Getenv("PORT")
	lineConfiguration := bot.LineConfiguration{
		LineToken: os.Getenv("LINE_TOKEN"),
		LineSecret: os.Getenv("LINE_SECRET"),
	}

	oxfordConfiguration := bot.OxfordConfiguration{
		AppID: os.Getenv("OXFORD_APPID"),
		AppKey: os.Getenv("OXFORD_APPKEY"),
	}

	app, err := bot.NewLineDictionaryApp(lineConfiguration, oxfordConfiguration)
	if err != nil {
		log.Fatal(err)
	}

	http.HandleFunc("/webhook", app.WebHook())
	log.Println("Start server on port:", port)
	if err := http.ListenAndServe(":" + port, nil); err != nil {
		log.Println(err)
	}
}