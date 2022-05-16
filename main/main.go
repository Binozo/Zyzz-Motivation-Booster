package main

import (
	"Zyzz-Motivation-Booster/pkg/cron"
	"Zyzz-Motivation-Booster/pkg/storage"
	"Zyzz-Motivation-Booster/pkg/telegram"
	"log"
	"os"
)

func main() {
	if err := storage.Load(); err != nil {
		panic(err)
	}

	if os.Getenv("telegrambottoken") == "" {
		panic("telegrambottoken env var is not set")
	}
	if os.Getenv("telegramchatid") == "" {
		panic("telegramchatid env var is not set")
	}

	log.Println("Starting...")
	telegram.SendMessage("Starting...")

	log.Println("Wating till 7 am...")
	cron.Setup()
}
