package template

import (
	"fmt"
	currencyCode "golang_telegram_bot/internal/enums/currency/code"
	"golang_telegram_bot/internal/models/currency"
)

func BuildTemplate(currencies map[string]*currency.Currency) string {
	return "Ежедневные данные о курсах валют на 11:30 по Мск(GMT+3)\n\n" +
		"💵 Курс $: " +
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
