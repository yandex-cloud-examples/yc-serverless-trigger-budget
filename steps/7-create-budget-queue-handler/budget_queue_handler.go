package main

import (
	"context"
	"fmt"
	"os"

	"github.com/pkg/errors"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func HandleBudgetQueueMessage(ctx context.Context, request interface{}) error {
	botAPIToken := os.Getenv("TELEGRAM_BOT_API_TOKEN")
	if botAPIToken == "" {
		return fmt.Errorf("TELEGRAM_BOT_API_TOKEN env variable is not set")
	}
	bot, err := tgbotapi.NewBotAPI(botAPIToken)
	if err != nil {
		return errors.Wrap(err, "initialize telegram bot")
	}

	for _, chatID := range getSubscribedChatIDs() {
		err = sendMessage(bot, chatID, "Budget trigger was triggered!")
		if err != nil {
			return errors.Wrap(err, "send message")
		}
	}

	return nil
}

func sendMessage(bot *tgbotapi.BotAPI, chatID int64, message string) error {
	msg := tgbotapi.NewMessage(chatID, message)
	msg.ParseMode = "Markdown"
	_, err := bot.Send(msg)
	return err
}

func getSubscribedChatIDs() []int64 {
	return []int64{requireEnvInt64("TELEGRAM_BOT_CHAT_ID")}
}
