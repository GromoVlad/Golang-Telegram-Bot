package currency

import (
	"fmt"
	"golang_telegram_bot/internal/DB"
	currencyCode "golang_telegram_bot/internal/enums/currency/code"
	"golang_telegram_bot/internal/models/currency"
	"golang_telegram_bot/internal/repository/currencyRepository"
	"log"
	"strconv"
	"strings"
)

func GetCurrencies() map[string]*currency.Currency {
	currencyTable := DB.Connect().Collection("currencies")
	allCurrency := currencyRepository.GetAllCurrency(currencyTable)
	currencies := make(map[string]*currency.Currency)
	for _, currency := range allCurrency {
		currencies[currency.Code] = currency
	}

	return currencies
}

func UpdateCurrencies() {
	currencies := currencyRepository.GetCBRFCurrency()

	for _, currencyNode := range currencies.Currency {
		var amount float64
		code := currencyNode.CharCode

		nominal, err := strconv.ParseFloat(strings.ReplaceAll(currencyNode.Nominal, ",", "."), 64)
		if err != nil {
			log.Fatalln(err)
		}
		value, err := strconv.ParseFloat(strings.ReplaceAll(currencyNode.Value, ",", "."), 64)
		if err != nil {
			log.Fatalln(err)
		}

		if code == currencyCode.USD || code == currencyCode.EUR || code == currencyCode.GBP ||
			code == currencyCode.CHF || code == currencyCode.BYN {
			amount = value
		} else if code == currencyCode.UAH || code == currencyCode.HKD || code == currencyCode.TJS ||
			code == currencyCode.CNY || code == currencyCode.TRY {
			amount = value / nominal
		} else if code == currencyCode.KGS || code == currencyCode.UZS || code == currencyCode.AMD ||
			code == currencyCode.INR || code == currencyCode.KZT {
			amount = nominal / value
		}

		if amount != 0.0 {
			currencyRepository.UpdateCurrency(code, amount)
		}
	}
}

func GetActualQuotes() string {
	actualCurrency := currencyRepository.GetActualCurrency()
	actualCryptocurrency := currencyRepository.GetActualCryptocurrency()

	return fmt.Sprintf(
		"🇺🇸 Курс $: %.2f₽\n"+
			"🇪🇺 Курс €: %.2f₽\n"+
			"🇰🇿 Курс ₸ (за ₽): %.2f₸\n"+
			"🇰🇿 Курс ₸ (за $): %.2f₸\n"+
			"🔶 Курс BTC: %.0f $\n"+
			"🔷 Курс ETH: %.0f $\n",
		actualCurrency.Quotes.USDRUB,
		actualCurrency.Quotes.USDRUB/actualCurrency.Quotes.USDEUR,
		actualCurrency.Quotes.USDKZT/actualCurrency.Quotes.USDRUB,
		actualCurrency.Quotes.USDKZT,
		actualCryptocurrency.Rates.BTC,
		actualCryptocurrency.Rates.ETH,
	)
}
