package sender

import (
	"fmt"
	"github.com/go-co-op/gocron"
	telegramBotAPI "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	currencyCode "golang_telegram_bot/internal/enums/currency/code"
	"golang_telegram_bot/internal/repository/telegramBotRepository"
	currencyService "golang_telegram_bot/internal/service/currency"
	"log"
	"os"
	"time"
)

func RunCronJobs() {
	scheduler := gocron.NewScheduler(time.UTC)
	currencies := currencyService.GetActualCurrencies()
	template := buildTemplate(currencies)

	scheduler.Every(1).Cron("32 08 * * *").Do(func() {
		headerTenHoursAmTemplate := "Ежедневные данные о курсах валют на 11:30 по Мск(GMT+3)\n\n" + template
		dispatch(headerTenHoursAmTemplate)
	})

	scheduler.StartAsync()
}

type Client struct {
	bot *telegramBotAPI.BotAPI
}

func New(apiKey string) *Client {
	bot, err := telegramBotAPI.NewBotAPI(apiKey)
	if err != nil {
		log.Panic(err)
	}

	return &Client{bot: bot}
}

func (c *Client) SendMessage(text string, chatId int64) error {
	msg := telegramBotAPI.NewMessage(chatId, text)
	msg.ParseMode = "Markdown"
	_, err := c.bot.Send(msg)
	return err
}

func dispatch(template string) {
	client := New(os.Getenv("TELEGRAM_TOKEN"))
	contacts := telegramBotRepository.FindSubscribers()

	for _, value := range contacts {
		err := client.SendMessage(template, int64(value.TelegramId))
		if err != nil {
			fmt.Println(err)
		}
	}
}

func buildTemplate(currencies map[string]currencyService.Currency) string {
	return "💵 Курс $: " +
		fmt.Sprintf("%.2f", currencies[currencyCode.USD].Amount) + currencies[currencyCode.USD].Icon +
		"\n💶 Курс €: " +
		fmt.Sprintf("%.2f", currencies[currencyCode.EUR].Amount) + currencies[currencyCode.EUR].Icon +
		"\n🇨🇭 Курс ₣: " +
		fmt.Sprintf("%.2f", currencies[currencyCode.CHF].Amount) + currencies[currencyCode.CHF].Icon +
		"\n🇬🇧 Курс £: " +
		fmt.Sprintf("%.2f", currencies[currencyCode.GBP].Amount) + currencies[currencyCode.GBP].Icon +
		"\n🇰🇿 Курс ₸ (за ₽): " +
		fmt.Sprintf("%.2f", currencies[currencyCode.KZT].Amount) + currencies[currencyCode.KZT].Icon +
		"\n🇹🇯 Курс сомони: " +
		fmt.Sprintf("%.2f", currencies[currencyCode.TJS].Amount) + currencies[currencyCode.TJS].Icon +
		"\n🇹🇷 Курс лиры: " +
		fmt.Sprintf("%.2f", currencies[currencyCode.TRY].Amount) + currencies[currencyCode.TRY].Icon +
		"\n🇦🇲 Курс драма (за ₽): " +
		fmt.Sprintf("%.2f", currencies[currencyCode.AMD].Amount) + currencies[currencyCode.AMD].Icon +
		"\n🇨🇳 Курс ¥: " +
		fmt.Sprintf("%.2f", currencies[currencyCode.CNY].Amount) + currencies[currencyCode.CNY].Icon +
		"\n🇭🇰 Курс HKD: " +
		fmt.Sprintf("%.2f", currencies[currencyCode.HKD].Amount) + currencies[currencyCode.HKD].Icon +
		"\n🇮🇳 Курс рупии (за ₽): " +
		fmt.Sprintf("%.2f", currencies[currencyCode.INR].Amount) + currencies[currencyCode.INR].Icon +
		"\n🇰🇬 Курс сома (за ₽): " +
		fmt.Sprintf("%.2f", currencies[currencyCode.KGS].Amount) + currencies[currencyCode.KGS].Icon +
		"\n🇺🇿 Курс сума (за ₽): " +
		fmt.Sprintf("%.2f", currencies[currencyCode.UZS].Amount) + currencies[currencyCode.UZS].Icon +
		"\n🇺🇦 Курс гривны: " +
		fmt.Sprintf("%.2f", currencies[currencyCode.UAH].Amount) + currencies[currencyCode.UAH].Icon
}
