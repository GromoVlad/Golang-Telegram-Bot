package DB

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"os"
)

var Context = context.TODO()
var Connect = func() *mongo.Database {
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
