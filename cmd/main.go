package main

import (
	"fmt"
	"github.com/arm01846/dictionary-bot/pkg"
	"io/ioutil"
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

	http.HandleFunc("/webhook", lineWebHookHandler(lineConfiguration))
	log.Println("Start server on port:", port)
	if err := http.ListenAndServe(":" + port, nil); err != nil {
		log.Println(err)
	}
}

func lineWebHookHandler(lineConfiguration bot.LineConfiguration) http.HandlerFunc {
	signatureValidator := bot.NewLineSignatureValidator(lineConfiguration.LineSecret)

	return func(w http.ResponseWriter, r *http.Request) {
		defer r.Body.Close()
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			log.Println(err)
			return
		}

		if ok := signatureValidator.Validate(r.Header.Get("X-Line-Signature"), body); !ok {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		fmt.Println(string(body))
	}
}