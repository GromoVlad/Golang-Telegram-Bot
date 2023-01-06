package listener

import (
	telegramBotAPI "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"golang_telegram_bot/internal/repository/telegramBotRepository"
	currencyService "golang_telegram_bot/internal/service/currency"
	"golang_telegram_bot/internal/support/telegramBotKeyboard"
	"log"
	"os"
)

const (
	CurrencyResponse  = "Актуальный курс валют ЦБ РФ 🏦"
	SubscribeResponse = "Вы успешно подписались на рассылку 📳\n\nСообщение будет отправляться ежедневно " +
		"в 11:30 (по Мск/GMT+3) 🕦"
	DescribeResponse = "Вы успешно отписались от рассылки 👌"
	DefaultResponse  = "Не удалось обработать сообщение 🤷🏻\nПопробуйте воспользоваться меню 👇🏻"
	StartResponse    = "Добро пожаловать! 👋🏻\n\nМы поможем Вам получить актуальный курс валют 💹\n" +
		"Информация основана на данных ЦБ РФ 🏦\n\nВыберите интересующий Вас вариант в меню 🧐"
)

func StartBot() {
	bot, err := telegramBotAPI.NewBotAPI(os.Getenv("TELEGRAM_TOKEN"))
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true

	log.Printf("Authorized on account %s\n", bot.Self.UserName)

	updateConfig := telegramBotAPI.NewUpdate(0)
	updateConfig.Timeout = 60
	updatesChannel := bot.GetUpdatesChan(updateConfig)

	// Циклически просматривайте каждое обновление
	for update := range updatesChannel {
		// Проверяем, получили ли мы сообщения
		if update.Message != nil {

			contact := telegramBotRepository.FindTelegramContact(update.Message)
			if contact.TelegramId == 0 {
				telegramBotRepository.InsertTelegramContact(update.Message)
			}

			// Создаем новое сообщение с id чата и текстом который мы получили
			message := telegramBotAPI.NewMessage(update.Message.Chat.ID, update.Message.Text)
			subscribeContact := telegramBotRepository.FindTelegramContact(update.Message)
			message.ReplyMarkup = telegramBotKeyboard.GetBaseKeyboard(subscribeContact.IsSubscriber)

			switch update.Message.Text {
			case telegramBotKeyboard.CURRENCY:
				message.Text = CurrencyResponse
				message.ReplyMarkup = telegramBotKeyboard.GetCurrenciesKeyboard(currencyService.GetActualCurrencies())

			case telegramBotKeyboard.SUBSCRIBE:
				telegramBotRepository.UpdateTelegramContact(int(update.Message.Chat.ID), true)
				message.Text = SubscribeResponse
				message.ReplyMarkup = telegramBotKeyboard.GetBaseKeyboard(true)

			case telegramBotKeyboard.DESCRIBE:
				telegramBotRepository.UpdateTelegramContact(int(update.Message.Chat.ID), false)
				message.Text = DescribeResponse
				message.ReplyMarkup = telegramBotKeyboard.GetBaseKeyboard(false)

			default:
				message.Text = DefaultResponse
			}

			switch update.Message.Command() {
			case "start":
				message.Text = StartResponse
			}

			// Отправляем сообщение
			if _, err = bot.Send(message); err != nil {
				panic(err)
			}

		} else if update.CallbackQuery != nil {
			// Ответьте на запрос callback-a, показав в Telegram пользователю сообщение с полученными данными
			callback := telegramBotAPI.NewCallback(update.CallbackQuery.ID, update.CallbackQuery.Data)
			if _, err := bot.Request(callback); err != nil {
				panic(err)
			}

			// И, наконец, отправьте сообщение, содержащее полученные данные
			msg := telegramBotAPI.NewMessage(update.CallbackQuery.Message.Chat.ID, update.CallbackQuery.Data)
			if _, err := bot.Send(msg); err != nil {
				panic(err)
			}
		}
	}
}
