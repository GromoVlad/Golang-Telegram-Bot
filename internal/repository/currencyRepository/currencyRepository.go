package currencyRepository

import (
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"golang_telegram_bot/internal/DB"
	"golang_telegram_bot/internal/models/currency"
	"golang_telegram_bot/internal/support/template"
)

func FindOneCurrency(code string) currency.Currency {
	currencyTable := DB.Connect().Collection("currencies")
	filter := bson.D{{"code", code}}
	return filterCurrency(currencyTable, filter)
}

func GetAllCurrency(currencyTable *mongo.Collection) []*currency.Currency {
	filter := bson.D{{}}
	currencies, _ := filterCurrencies(currencyTable, filter)

	fmt.Println("--------------GetAllCurrency---------------")
	for _, currency := range currencies {
		fmt.Printf(
			"id = %v |  Amount = %v | Icon = %v | Name = %v | Code = %v | CreatedAt = %v \n",
			currency.ID, currency.Amount, currency.Icon, currency.Name, currency.Code, currency.CreatedAt,
		)
	}

	return currencies
}

func UpdateCurrency(code string, amount float64) {
	currencyTable := DB.Connect().Collection("currencies")
	currency := FindOneCurrency(code)
	currency.Amount = amount
	update := bson.M{"$set": currency}
	filter := bson.D{{"code", code}}
	_, err := currencyTable.UpdateOne(DB.Context, filter, update)
	if err != nil {
		fmt.Println(err)
	}
}

func StartSeedCurrency() map[string]*currency.Currency {
	currencyTable := DB.Connect().Collection("currencies")

	filter := bson.D{{}}
	currencyTable.DeleteMany(DB.Context, filter)
	currencies := template.BaseCurrenciesTemplate()

	for _, currency := range currencies {
		currencyTable.InsertOne(DB.Context, currency)
	}

	return currencies
}

func filterCurrency(currencyTable *mongo.Collection, filter interface{}) currency.Currency {
	var currency currency.Currency
	singleResult := currencyTable.FindOne(DB.Context, filter)
	singleResult.Decode(&currency)
	return currency
}

func filterCurrencies(collection *mongo.Collection, filter interface{}) ([]*currency.Currency, error) {
	var currencies []*currency.Currency

	cursor, err := collection.Find(DB.Context, filter)
	if err != nil {
		return currencies, err
	}

	for cursor.Next(DB.Context) {
		var currency currency.Currency
		err := cursor.Decode(&currency)
		if err != nil {
			return currencies, err
		}

		currencies = append(currencies, &currency)
	}

	if err := cursor.Err(); err != nil {
		return currencies, err
	}

	cursor.Close(DB.Context)

	if len(currencies) == 0 {
		return currencies, mongo.ErrNoDocuments
	}

	return currencies, nil
}
