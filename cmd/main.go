package main

import (
	"golang_telegram_bot/internal/jobs"
	"golang_telegram_bot/internal/repository/currencyRepository"
	currencyService "golang_telegram_bot/internal/service/currency"
	"golang_telegram_bot/internal/service/telegramBot"
	"time"
)

func main() {
	// Запускаем стартовый seeder с шаблоном курсов валют
	currencyRepository.StartSeedCurrency()
	// Обновляем данные о валютах после seed-а
	currencyService.UpdateCurrencies()

	// Старт служб Telegram бота
	go telegramBot.StartBot()
	go jobs.RunCronJobs()

	// Спим, пока работают горутины
	time.Sleep(365 * 24 * time.Hour)
}
