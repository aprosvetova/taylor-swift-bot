package main

import (
	tb "gopkg.in/tucnak/telebot.v2"
	"log"
	"os"
	"time"
)

var tg *tb.Bot

func main() {
	var err error

	tg, err = tb.NewBot(tb.Settings{
		Token:  os.Getenv("TG_TOKEN"),
		Poller: &tb.LongPoller{Timeout: 10 * time.Second},
	})

	if err != nil {
		log.Fatalln("Telegram", err)
		return
	}

	tg.Handle("/start", handleTaylor)
	tg.Handle("/taylor", handleTaylor)

	tg.Start()
}

func handleTaylor(m *tb.Message) {
	img, quote, err := getTaylor()
	if err != nil {
		tg.Reply(m, "Error: "+err.Error())
		return
	}

	if img == "" {
		tg.Reply(m, quote)
		return
	}

	tg.Reply(m, &tb.Photo{
		File:    tb.FromURL(img),
		Caption: quote,
	})
}
