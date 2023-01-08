package userRepository

import (
	"fmt"
	telegramBotAPI "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"golang_telegram_bot/internal/DB"
	. "golang_telegram_bot/internal/models/user"
	"time"
)

func GetAllUsers() {
	filter := bson.D{{}}
	users, _ := filterContacts(filter)

	fmt.Println("--------------GetAllUsers---------------")
	for _, user := range users {
		fmt.Printf(
			"id = %v |  TelegramId = %v | TelegramName = %v | IsSubscriber = %v | CreatedAt = %v \n",
			user.ID, user.TelegramId, user.TelegramName, user.IsSubscriber, user.CreatedAt,
		)
	}
}

func FindOneTelegramContact(telegramId int) UserContact {
	filter := bson.D{{"telegram_id", telegramId}}
	result, _ := filterContact(filter)
	return result
}

func filterContacts(filter interface{}) ([]*UserContact, error) {
	userTable := DB.Connect().Collection("user")
	var contacts []*UserContact

	cursor, err := userTable.Find(DB.Context, filter)
	if err != nil {
		return contacts, err
	}

	for cursor.Next(DB.Context) {
		var contact UserContact
		err := cursor.Decode(&contact)
		if err != nil {
			return contacts, err
		}

		contacts = append(contacts, &contact)
	}

	if err := cursor.Err(); err != nil {
		return contacts, err
	}

	cursor.Close(DB.Context)

	if len(contacts) == 0 {
		return contacts, mongo.ErrNoDocuments
	}

	return contacts, nil
}

func filterContact(filter interface{}) (UserContact, error) {
	userTable := DB.Connect().Collection("user")
	var contact UserContact
	singleResult := userTable.FindOne(DB.Context, filter)
	singleResult.Decode(&contact)
	return contact, nil
}

func InsertTelegramContact(messageData *telegramBotAPI.Message) {
	userTable := DB.Connect().Collection("user")
	telegramName := messageData.Chat.UserName
	if telegramName == "" {
		telegramName = messageData.Chat.FirstName + " " + messageData.Chat.LastName
	}

	contact := &UserContact{
		ID:           primitive.NewObjectID(),
		TelegramId:   int(messageData.Chat.ID),
		TelegramName: telegramName,
		IsSubscriber: false,
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	}

	userTable.InsertOne(DB.Context, contact)
}

func UpdateTelegramContact(telegramId int, isSubscriber bool) {
	userTable := DB.Connect().Collection("user")
	contact := FindOneTelegramContact(telegramId)
	contact.IsSubscriber = isSubscriber
	update := bson.M{"$set": contact}
	filter := bson.D{{"telegram_id", telegramId}}

	_, err := userTable.UpdateOne(DB.Context, filter, update)
	if err != nil {
		fmt.Println(err)
	}
}

func FindSubscribers() []*UserContact {
	DB.Connect().Collection("user")
	filter := bson.D{{"is_subscriber", true}}
	contacts, _ := filterContacts(filter)

	return contacts
}
