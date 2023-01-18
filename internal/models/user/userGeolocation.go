package userContact

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type UserGeolocation struct {
	ID         primitive.ObjectID `bson:"id"`
	UserId     int                `bson:"user_id"`
	Longitude  float64            `bson:"longitude"`
	Latitude   float64            `bson:"latitude"`
	NeedUpdate bool               `bson:"need_update"`
	CreatedAt  time.Time          `bson:"created_at"`
	UpdatedAt  time.Time          `bson:"updated_at"`
}
