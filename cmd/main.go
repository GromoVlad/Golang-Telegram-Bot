package main

import (
	"github.com/joho/godotenv"
	telegramBotListener "golang_telegram_bot/internal/service/telegramBot/listener"
	telegramBotSender "golang_telegram_bot/internal/service/telegramBot/sender"
	"log"
)

func main() {
	/** Подгрудаем данные из .env */
	if err := godotenv.Load(); err != nil {
		log.Fatalf("error loading env variables: %s", err.Error())
	}

	/** Старт служб Telegram бота */
	telegramBotListener.StartBot()
	telegramBotSender.RunCronJobs()
}
