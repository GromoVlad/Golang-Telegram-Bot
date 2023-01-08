package DB

import (
	"context"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"os"
)

var Context = context.TODO()
var Connect = func() *mongo.Database {
	// Подгружаем данные из .env
	if err := godotenv.Load(); err != nil {
		log.Printf("Ошибка загрузки переменных env: %s", err.Error())
	}

	clientOptions := options.Client().ApplyURI(os.Getenv("MONGODB_CONNECTION"))
	client, err := mongo.Connect(Context, clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(Context, nil)
	if err != nil {
		log.Fatal(err)
	}

	return client.Database("currency_telegram_bot")
}
