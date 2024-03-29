package telegramBotKeyboard

import (
	"fmt"
	telegramBotAPI "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	currencyCode "golang_telegram_bot/internal/enums/currency/code"
	"golang_telegram_bot/internal/models/currency"
)

const (
	ACTUAL_CURRENCY = "Получить актуальные курсы валют 📈"
	CURRENCY        = "Получить курсы валют ЦБ РФ 🏦"
	DESCRIBE        = "Отписаться от рассылки курса валют ❌"
	SUBSCRIBE       = "Подписаться на рассылку курса валют 💬"
	WEATHER         = "Какая сейчас погода? 🏖"
	GEO             = "Я хочу обновить свои геоданные! 🧭"
)

func GetBaseKeyboard(isSubscribe bool, isUpdateGeolocation bool) telegramBotAPI.ReplyKeyboardMarkup {
	weatherButton := telegramBotAPI.NewKeyboardButton(WEATHER)
	weatherButton.RequestLocation = isUpdateGeolocation

	if isSubscribe {
		return telegramBotAPI.NewReplyKeyboard(
			telegramBotAPI.NewKeyboardButtonRow(telegramBotAPI.NewKeyboardButton(ACTUAL_CURRENCY)),
			telegramBotAPI.NewKeyboardButtonRow(telegramBotAPI.NewKeyboardButton(CURRENCY)),
			telegramBotAPI.NewKeyboardButtonRow(telegramBotAPI.NewKeyboardButton(DESCRIBE)),
			telegramBotAPI.NewKeyboardButtonRow(weatherButton),
			telegramBotAPI.NewKeyboardButtonRow(telegramBotAPI.NewKeyboardButton(GEO)),
		)
	} else {
		return telegramBotAPI.NewReplyKeyboard(
			telegramBotAPI.NewKeyboardButtonRow(telegramBotAPI.NewKeyboardButton(ACTUAL_CURRENCY)),
			telegramBotAPI.NewKeyboardButtonRow(telegramBotAPI.NewKeyboardButton(CURRENCY)),
			telegramBotAPI.NewKeyboardButtonRow(telegramBotAPI.NewKeyboardButton(SUBSCRIBE)),
			telegramBotAPI.NewKeyboardButtonRow(weatherButton),
			telegramBotAPI.NewKeyboardButtonRow(telegramBotAPI.NewKeyboardButton(GEO)),
		)
	}
}

func GetCurrenciesKeyboard(currencies map[string]*currency.Currency) telegramBotAPI.InlineKeyboardMarkup {
	return telegramBotAPI.NewInlineKeyboardMarkup(
		telegramBotAPI.NewInlineKeyboardRow(
			telegramBotAPI.NewInlineKeyboardButtonData(
				"💵 Доллар США (в ₽)",
				"💵 Курс $: "+
					fmt.Sprintf("%.2f", currencies[currencyCode.USD].Amount)+currencies[currencyCode.USD].Icon,
			),
			telegramBotAPI.NewInlineKeyboardButtonData(
				"💶 Евро (в ₽)",
				"💶 Курс €: "+
					fmt.Sprintf("%.2f", currencies[currencyCode.EUR].Amount)+currencies[currencyCode.EUR].Icon,
			),
		),
		telegramBotAPI.NewInlineKeyboardRow(
			telegramBotAPI.NewInlineKeyboardButtonData(
				"🇨🇭 Швейцарский франк (в ₽)",
				"🇨🇭 Курс ₣: "+
					fmt.Sprintf("%.2f", currencies[currencyCode.CHF].Amount)+currencies[currencyCode.CHF].Icon,
			),
			telegramBotAPI.NewInlineKeyboardButtonData(
				"🇬🇧 Фунт стерлингов (в ₽)",
				"🇬🇧 Курс £: "+
					fmt.Sprintf("%.2f", currencies[currencyCode.GBP].Amount)+currencies[currencyCode.GBP].Icon,
			),
		),
		telegramBotAPI.NewInlineKeyboardRow(
			telegramBotAPI.NewInlineKeyboardButtonData(
				"🇰🇿 Казахстанский тенге (за ₽)",
				"🇰🇿 Курс ₸ (за ₽): "+
					fmt.Sprintf("%.2f", currencies[currencyCode.KZT].Amount)+currencies[currencyCode.KZT].Icon,
			),
			telegramBotAPI.NewInlineKeyboardButtonData(
				"🇹🇯 Таджикский сомони (в ₽)",
				"🇹🇯 Курс сомони: "+
					fmt.Sprintf("%.2f", currencies[currencyCode.TJS].Amount)+currencies[currencyCode.TJS].Icon,
			),
		),
		telegramBotAPI.NewInlineKeyboardRow(
			telegramBotAPI.NewInlineKeyboardButtonData(
				"🇹🇷 Турецкая лира (в ₽)",
				"🇹🇷 Курс лиры: "+
					fmt.Sprintf("%.2f", currencies[currencyCode.TRY].Amount)+currencies[currencyCode.TRY].Icon,
			),
			telegramBotAPI.NewInlineKeyboardButtonData(
				"🇦🇲 Армянский драм (за ₽)",
				"🇦🇲 Курс драма (за ₽): "+
					fmt.Sprintf("%.2f", currencies[currencyCode.AMD].Amount)+currencies[currencyCode.AMD].Icon,
			),
		),
		telegramBotAPI.NewInlineKeyboardRow(
			telegramBotAPI.NewInlineKeyboardButtonData(
				"🇨🇳 Китайский юань (в ₽)",
				"🇨🇳 Курс ¥: "+
					fmt.Sprintf("%.2f", currencies[currencyCode.CNY].Amount)+currencies[currencyCode.CNY].Icon,
			),
			telegramBotAPI.NewInlineKeyboardButtonData(
				"🇭🇰 Гонконгский доллар (в ₽)",
				"🇭🇰 Курс HKD: "+
					fmt.Sprintf("%.2f", currencies[currencyCode.HKD].Amount)+currencies[currencyCode.HKD].Icon,
			),
		),
		telegramBotAPI.NewInlineKeyboardRow(
			telegramBotAPI.NewInlineKeyboardButtonData(
				"🇮🇳 Индийская рупия (за ₽)",
				"🇮🇳 Курс рупии (за ₽): "+
					fmt.Sprintf("%.2f", currencies[currencyCode.INR].Amount)+currencies[currencyCode.INR].Icon,
			),
			telegramBotAPI.NewInlineKeyboardButtonData(
				"🇰🇬 Киргизский сом (за ₽)",
				"🇰🇬 Курс сома (за ₽): "+
					fmt.Sprintf("%.2f", currencies[currencyCode.KGS].Amount)+currencies[currencyCode.KGS].Icon,
			),
		),
		telegramBotAPI.NewInlineKeyboardRow(
			telegramBotAPI.NewInlineKeyboardButtonData(
				"🇺🇿 Узбекский сум (в ₽)",
				"🇺🇿 Курс сума (за ₽): "+
					fmt.Sprintf("%.2f", currencies[currencyCode.UZS].Amount)+currencies[currencyCode.UZS].Icon,
			),
			telegramBotAPI.NewInlineKeyboardButtonData(
				"🇺🇦 Украинская гривна (в ₽)",
				"🇺🇦 Курс гривны: "+
					fmt.Sprintf("%.2f", currencies[currencyCode.UAH].Amount)+currencies[currencyCode.UAH].Icon,
			),
		),
	)
}
