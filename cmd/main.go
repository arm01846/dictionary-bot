package main

import (
	"github.com/arm01846/dictionary-bot/pkg"
	"log"
	"net/http"
	"os"
)

func main() {
	var port = os.Getenv("PORT")
	var lineConfiguration = bot.LineConfiguration{
		LineToken: os.Getenv("LINE_TOKEN"),
		LineSecret: os.Getenv("LINE_SECRET"),
	}

	app, err := bot.NewLineDictionaryApp(lineConfiguration)
	if err != nil {
		log.Fatal(err)
	}

	http.HandleFunc("/webhook", app.WebHook())
	log.Println("Start server on port:", port)
	if err := http.ListenAndServe(":" + port, nil); err != nil {
		log.Println(err)
	}
}