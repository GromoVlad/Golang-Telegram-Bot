package telegramBotKeyboard

import (
	"fmt"
	telegramBotAPI "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	currencyCode "golang_telegram_bot/internal/enums/currency/code"
	currencyService "golang_telegram_bot/internal/service/currency"
)

const (
	CURRENCY  = "Получить курсы валют ЦБ РФ 💵 📈"
	DESCRIBE  = "Отписаться от рассылки курса валют ❌"
	SUBSCRIBE = "Подписаться на рассылку курса валют 💬"
)

func GetBaseKeyboard(isSubscribe bool) telegramBotAPI.ReplyKeyboardMarkup {
	if isSubscribe {
		return telegramBotAPI.NewReplyKeyboard(
			telegramBotAPI.NewKeyboardButtonRow(telegramBotAPI.NewKeyboardButton(CURRENCY)),
			telegramBotAPI.NewKeyboardButtonRow(telegramBotAPI.NewKeyboardButton(DESCRIBE)),
		)
	} else {
		return telegramBotAPI.NewReplyKeyboard(
			telegramBotAPI.NewKeyboardButtonRow(telegramBotAPI.NewKeyboardButton(CURRENCY)),
			telegramBotAPI.NewKeyboardButtonRow(telegramBotAPI.NewKeyboardButton(SUBSCRIBE)),
		)
	}
}

func GetCurrenciesKeyboard(currencies map[string]currencyService.Currency) telegramBotAPI.InlineKeyboardMarkup {
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
