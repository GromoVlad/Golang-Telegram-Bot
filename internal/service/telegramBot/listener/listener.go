package listener

import (
	telegramBotAPI "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"golang_telegram_bot/internal/repository/userRepository"
	currencyService "golang_telegram_bot/internal/service/currency"
	"golang_telegram_bot/internal/support/telegramBotKeyboard"
	"golang_telegram_bot/internal/support/template"
)

func MessageHandler(messageData *telegramBotAPI.Message, bot *telegramBotAPI.BotAPI) {

	userRepository.GetAllUsers()

	contact := userRepository.FindOneTelegramContact(int(messageData.Chat.ID))
	if contact.TelegramId == 0 {
		userRepository.InsertTelegramContact(messageData)
	}

	message := telegramBotAPI.NewMessage(messageData.Chat.ID, messageData.Text)
	subscribeContact := userRepository.FindOneTelegramContact(int(messageData.Chat.ID))
	message.ReplyMarkup = telegramBotKeyboard.GetBaseKeyboard(subscribeContact.IsSubscriber)

	switch messageData.Text {
	case telegramBotKeyboard.CURRENCY:
		message.Text = template.CurrencyResponse
		currencies := currencyService.GetCurrencies()
		message.ReplyMarkup = telegramBotKeyboard.GetCurrenciesKeyboard(currencies)

	case telegramBotKeyboard.SUBSCRIBE:
		userRepository.UpdateTelegramContact(int(messageData.Chat.ID), true)
		message.Text = template.SubscribeResponse
		message.ReplyMarkup = telegramBotKeyboard.GetBaseKeyboard(true)

	case telegramBotKeyboard.DESCRIBE:
		userRepository.UpdateTelegramContact(int(messageData.Chat.ID), false)
		message.Text = template.DescribeResponse
		message.ReplyMarkup = telegramBotKeyboard.GetBaseKeyboard(false)

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
