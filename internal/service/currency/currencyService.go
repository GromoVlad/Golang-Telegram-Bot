package currency

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"golang.org/x/net/html/charset"
	currencyCode "golang_telegram_bot/internal/enums/currency/code"
	currencyIcon "golang_telegram_bot/internal/enums/currency/icon"
	"io"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"
)

const CBRF_URL = "https://www.cbr.ru/scripts/XML_daily.asp?date_req="

type CurrenciesCBRF struct {
	XMLName  xml.Name       `xml:"ValCurs"`
	Currency []CurrencyCBRF `xml:"Valute"`
}

type CurrencyCBRF struct {
	NumCode  string `xml:"NumCode"`
	CharCode string `xml:"CharCode"`
	Nominal  string `xml:"Nominal"`
	Name     string `xml:"Name"`
	Value    string `xml:"Value"`
}

type Currency struct {
	Amount float64
	Icon   string
	Name   string
	Code   string
}

func GetActualCurrencies() map[string]Currency {
	resp, err := http.Get(CBRF_URL + time.Now().Format("02/01/2006"))
	if err != nil {
		fmt.Println(err)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
	}

	currencies := new(CurrenciesCBRF)
	reader := bytes.NewReader(body)
	decoder := xml.NewDecoder(reader)
	decoder.CharsetReader = charset.NewReaderLabel
	err = decoder.Decode(&currencies)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("%v\n\n", currencies)

	necessaryCurrencies := make(map[string]Currency)
	for _, currencyNode := range currencies.Currency {

		var amount float64
		var icon string

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
			icon = currencyIcon.RUB
		} else if code == currencyCode.UAH || code == currencyCode.HKD || code == currencyCode.TJS ||
			code == currencyCode.CNY || code == currencyCode.TRY {
			amount = value / nominal
			icon = currencyIcon.RUB
		} else if code == currencyCode.KGS {
			amount = nominal / value
			icon = currencyIcon.KGS
		} else if code == currencyCode.UZS {
			amount = nominal / value
			icon = currencyIcon.UZS
		} else if code == currencyCode.AMD {
			amount = nominal / value
			icon = currencyIcon.AMD
		} else if code == currencyCode.INR {
			amount = nominal / value
			icon = currencyIcon.INR
		} else if code == currencyCode.KZT {
			amount = nominal / value
			icon = currencyIcon.KZT
		}

		if amount != 0.0 && icon != "" {
			data := Currency{
				Amount: amount,
				Icon:   icon,
				Name:   currencyNode.Name,
				Code:   code,
			}
			necessaryCurrencies[code] = data
		}
	}

	return necessaryCurrencies
}
