package main

import (
	"fmt"
	"log"
	"os"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func main() {
	botAPIToken := os.Getenv("TELEGRAM_BOT_API_TOKEN")
	if botAPIToken == "" {
		log.Fatal("TELEGRAM_BOT_API_TOKEN env variable is not set")
	}
	bot, err := tgbotapi.NewBotAPI(botAPIToken)
	if err != nil {
		log.Fatal(err)
	}

	updateConfig := tgbotapi.NewUpdate(0)
	updateConfig.Timeout = 60

	updates := bot.GetUpdatesChan(updateConfig)

	for update := range updates {
		if update.Message != nil {
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, fmt.Sprintf("%v", update.Message.Chat.ID))
			bot.Send(msg)
		}
	}
}
