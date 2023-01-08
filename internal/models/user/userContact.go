package userContact

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type UserContact struct {
	ID           primitive.ObjectID `bson:"id"`
	TelegramId   int                `bson:"telegram_id"`
	TelegramName string             `bson:"telegram_name"`
	IsSubscriber bool               `bson:"is_subscriber"`
	CreatedAt    time.Time          `bson:"created_at"`
	UpdatedAt    time.Time          `bson:"updated_at"`
}
