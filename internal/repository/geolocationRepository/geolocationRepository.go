package geolocationRepository

import (
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"golang_telegram_bot/internal/DB"
	userContact "golang_telegram_bot/internal/models/user"
	"time"
)

func FindGeolocation(userId int) userContact.UserGeolocation {
	geolocationTable := DB.Connect().Collection("geolocations")
	filter := bson.D{{"user_id", userId}}
	return filterGeolocation(geolocationTable, filter)
}

func GetAllGeolocation() {
	filter := bson.D{{}}
	geolocations, _ := filterGeolocations(filter)

	fmt.Println("--------------GetAllGeolocation---------------")
	for _, geolocation := range geolocations {
		fmt.Printf(
			"id = %v |  UserId = %v | NeedUpdate = %v | Longitude = %v | Latitude = %v | CreatedAt = %v \n",
			geolocation.ID, geolocation.UserId, geolocation.NeedUpdate, geolocation.Longitude, geolocation.Latitude, geolocation.CreatedAt,
		)
	}
}

func InsertGeolocation(userId int, longitude float64, latitude float64) userContact.UserGeolocation {
	geolocationTable := DB.Connect().Collection("geolocations")

	//filter := bson.D{{}}
	//geolocationTable.DeleteMany(DB.Context, filter)
	geolocation := userContact.UserGeolocation{
		ID:         primitive.NewObjectID(),
		UserId:     userId,
		Longitude:  longitude,
		Latitude:   latitude,
		NeedUpdate: false,
		CreatedAt:  time.Now(),
		UpdatedAt:  time.Now(),
	}

	geolocationTable.InsertOne(DB.Context, geolocation)

	return geolocation
}

func UpdateGeolocation(userId int, longitude float64, latitude float64, needUpdate bool) {
	geolocationTable := DB.Connect().Collection("geolocations")
	geolocation := FindGeolocation(userId)
	geolocation.Latitude = latitude
	geolocation.Longitude = longitude
	geolocation.NeedUpdate = needUpdate
	update := bson.M{"$set": geolocation}
	filter := bson.D{{"user_id", userId}}
	result, err := geolocationTable.UpdateOne(DB.Context, filter, update)

	fmt.Println("--------------UpdateGeolocation---------------")
	fmt.Println(result)
	if err != nil {
		fmt.Println(err)
	}
}

func filterGeolocation(geolocationTable *mongo.Collection, filter any) userContact.UserGeolocation {
	var geolocation userContact.UserGeolocation
	singleResult := geolocationTable.FindOne(DB.Context, filter)
	singleResult.Decode(&geolocation)
	return geolocation
}

func filterGeolocations(filter interface{}) ([]*userContact.UserGeolocation, error) {
	geolocationTable := DB.Connect().Collection("geolocations")
	var geolocations []*userContact.UserGeolocation

	cursor, err := geolocationTable.Find(DB.Context, filter)
	if err != nil {
		return geolocations, err
	}

	for cursor.Next(DB.Context) {
		var geolocation userContact.UserGeolocation
		err := cursor.Decode(&geolocation)
		if err != nil {
			return geolocations, err
		}

		geolocations = append(geolocations, &geolocation)
	}

	if err := cursor.Err(); err != nil {
		return geolocations, err
	}

	cursor.Close(DB.Context)

	if len(geolocations) == 0 {
		return geolocations, mongo.ErrNoDocuments
	}

	return geolocations, nil
}
