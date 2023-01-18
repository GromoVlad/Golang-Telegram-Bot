package telegramBot

import (
	telegramBotAPI "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"golang_telegram_bot/internal/service/telegramBot/listener"
	"log"
)

func StartBot() {
	bot, err := telegramBotAPI.NewBotAPI("5682843715:AAGoZweYYwlde7doz7gmnTuq3_SQNUVwllY")
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true
	log.Printf("Авторизован на учетной записи %s\n", bot.Self.UserName)
	updateConfig := telegramBotAPI.NewUpdate(0)
	updateConfig.Timeout = 60
	updatesChannel := bot.GetUpdatesChan(updateConfig)

	for update := range updatesChannel {
		if update.Message != nil {
			listener.MessageHandler(update.Message, bot)
		} else if update.CallbackQuery != nil {
			listener.CallbackHandler(update.CallbackQuery, bot)
		}
	}
}
