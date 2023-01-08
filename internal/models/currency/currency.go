package currency

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type Currency struct {
	ID        primitive.ObjectID `bson:"id"`
	Amount    float64            `bson:"amount"`
	Icon      string             `bson:"icon"`
	Name      string             `bson:"name"`
	Code      string             `bson:"code"`
	CreatedAt time.Time          `bson:"created_at"`
	UpdatedAt time.Time          `bson:"updated_at"`
}
