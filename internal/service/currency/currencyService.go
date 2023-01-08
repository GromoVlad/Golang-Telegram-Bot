package currency

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"golang.org/x/net/html/charset"
	"golang_telegram_bot/internal/DB"
	currencyCode "golang_telegram_bot/internal/enums/currency/code"
	"golang_telegram_bot/internal/models/currency"
	"golang_telegram_bot/internal/models/currency/xmlCBRF"
	"golang_telegram_bot/internal/repository/currencyRepository"
	"io"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"
)

const CBRF_URL = "https://www.cbr.ru/scripts/XML_daily.asp?date_req="

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
	response, err := http.Get(CBRF_URL + time.Now().Format("02/01/2006"))
	if err != nil {
		fmt.Println(err)
	}

	body, err := io.ReadAll(response.Body)
	if err != nil {
		fmt.Println(err)
	}

	currencies := new(xmlCBRF.Currencies)
	reader := bytes.NewReader(body)
	decoder := xml.NewDecoder(reader)
	decoder.CharsetReader = charset.NewReaderLabel
	err = decoder.Decode(&currencies)
	if err != nil {
		fmt.Println(err)
	}

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
