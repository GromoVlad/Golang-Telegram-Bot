package listener

import (
	telegramBotAPI "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	userContact "golang_telegram_bot/internal/models/user"
	"golang_telegram_bot/internal/repository/geolocationRepository"
	"golang_telegram_bot/internal/repository/userRepository"
	currencyService "golang_telegram_bot/internal/service/currency"
	weatherService "golang_telegram_bot/internal/service/weather"
	"golang_telegram_bot/internal/support/telegramBotKeyboard"
	"golang_telegram_bot/internal/support/template"
)

func MessageHandler(messageData *telegramBotAPI.Message, bot *telegramBotAPI.BotAPI) {

	userRepository.GetAllUsers()
	geolocationRepository.GetAllGeolocation()

	contact := userRepository.FindOneTelegramContact(int(messageData.Chat.ID))
	if contact.TelegramId == 0 {
		userRepository.InsertTelegramContact(messageData)
	}

	geolocation := geolocationRepository.FindGeolocation(int(messageData.Chat.ID))
	message := telegramBotAPI.NewMessage(messageData.Chat.ID, messageData.Text)
	subscribeContact := userRepository.FindOneTelegramContact(int(messageData.Chat.ID))
	message.ReplyMarkup = telegramBotKeyboard.GetBaseKeyboard(subscribeContact.IsSubscriber, geolocation.NeedUpdate)

	switch {
	case messageData.Text == telegramBotKeyboard.CURRENCY:
		message.Text = template.CurrencyResponse
		currencies := currencyService.GetCurrencies()
		message.ReplyMarkup = telegramBotKeyboard.GetCurrenciesKeyboard(currencies)

	case messageData.Text == telegramBotKeyboard.SUBSCRIBE:
		userRepository.UpdateTelegramContact(int(messageData.Chat.ID), true)
		message.Text = template.SubscribeResponse
		message.ReplyMarkup = telegramBotKeyboard.GetBaseKeyboard(true, geolocation.NeedUpdate)

	case messageData.Text == telegramBotKeyboard.DESCRIBE:
		userRepository.UpdateTelegramContact(int(messageData.Chat.ID), false)
		message.Text = template.DescribeResponse
		message.ReplyMarkup = telegramBotKeyboard.GetBaseKeyboard(false, geolocation.NeedUpdate)

	case messageData.Text == telegramBotKeyboard.GEO:
		geolocationRepository.UpdateGeolocation(geolocation.UserId, geolocation.Longitude, geolocation.Latitude, true)
		message.ReplyMarkup = telegramBotKeyboard.GetBaseKeyboard(contact.IsSubscriber, true)
		message.Text = "Геометка сброшена! 🚫\nЗапросите данные о погоде для обновления 🌐"

	case messageData.Text == telegramBotKeyboard.WEATHER:
		weather := weatherService.GetWeatherByGeo(geolocation.Latitude, geolocation.Longitude)
		message.Text = template.BuildWeatherTemplate(weather)

	case messageData.Text == telegramBotKeyboard.ACTUAL_CURRENCY:
		message.Text = currencyService.GetActualQuotes()

	case messageData.Location != nil:
		createOrUpdateGeolocation(geolocation, messageData)
		weather := weatherService.GetWeatherByGeo(messageData.Location.Latitude, messageData.Location.Longitude)
		message.Text = template.BuildWeatherTemplate(weather)

	default:
		message.Text = template.DefaultResponse
	}

	switch messageData.Command() {
	case "start":
		message.Text = template.StartResponse
	}

	if _, err := bot.Send(message); err != nil {
		panic(err)
	}
}

func CallbackHandler(callbackQuery *telegramBotAPI.CallbackQuery, bot *telegramBotAPI.BotAPI) {
	callback := telegramBotAPI.NewCallback(callbackQuery.ID, callbackQuery.Data)
	if _, err := bot.Request(callback); err != nil {
		panic(err)
	}

	msg := telegramBotAPI.NewMessage(callbackQuery.Message.Chat.ID, callbackQuery.Data)
	if _, err := bot.Send(msg); err != nil {
		panic(err)
	}
}

func createOrUpdateGeolocation(geolocation userContact.UserGeolocation, messageData *telegramBotAPI.Message) {
	if geolocation.UserId == 0 {
		geolocationRepository.InsertGeolocation(
			int(messageData.Chat.ID),
			messageData.Location.Longitude,
			messageData.Location.Latitude,
		)
	} else if geolocation.UserId != 0 && geolocation.NeedUpdate {
		geolocationRepository.UpdateGeolocation(
			geolocation.UserId,
			geolocation.Longitude,
			geolocation.Latitude,
			false,
		)
	}
}
