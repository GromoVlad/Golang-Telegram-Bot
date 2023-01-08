package jobs

import (
	"github.com/go-co-op/gocron"
	currencyService "golang_telegram_bot/internal/service/currency"
	"golang_telegram_bot/internal/service/sender"
	templateForMailing "golang_telegram_bot/internal/support/template"
	"time"
)

func RunCronJobs() {
	// Стартуем планировщик
	scheduler := gocron.NewScheduler(time.UTC)

	// Запускаем планировщик для обновления данных курса валют с сайта ЦБ РФ
	scheduler.Every(1).Cron("31 06 * * *").Do(func() {
		currencyService.UpdateCurrencies()
	})

	// Запускаем рассылку курса валют для подписчиков
	scheduler.Every(1).Cron("32 06 * * *").Do(func() {
		currencies := currencyService.GetCurrencies()
		template := templateForMailing.BuildTemplate(currencies)
		sender.Dispatch(template)
	})

	scheduler.StartBlocking()
}
