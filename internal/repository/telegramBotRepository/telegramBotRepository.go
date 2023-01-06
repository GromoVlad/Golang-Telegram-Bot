package telegramBotRepository

import (
	"database/sql"
	"fmt"
	telegramBotAPI "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"golang_telegram_bot/internal/database/DB"
	telegramContact "golang_telegram_bot/internal/models/contact/telegram"
	"time"
)

func FindTelegramContact(messageData *telegramBotAPI.Message) telegramContact.TelegramContact {
	var contact telegramContact.TelegramContact
	_ = DB.Connect().Get(
		&contact,
		"SELECT telegram_id, is_subscriber FROM contact.telegram WHERE telegram_id = $1",
		int(messageData.Chat.ID),
	)
	return contact
}

func InsertTelegramContact(messageData *telegramBotAPI.Message) {
	name := messageData.Chat.UserName
	if name == "" {
		name = messageData.Chat.FirstName + " " + messageData.Chat.LastName
	}

	transaction := DB.Connect().MustBegin()
	transaction.NamedExec(
		"INSERT INTO contact.telegram (telegram_id, telegram_name, created_at, updated_at) "+
			"VALUES (:telegram_id, :telegram_name, :created_at, :updated_at)",
		&telegramContact.TelegramContact{
			TelegramId:   int(messageData.Chat.ID),
			TelegramName: name,
			IsSubscriber: false,
			CreatedAt:    sql.NullTime{Time: time.Now(), Valid: true},
			UpdatedAt:    sql.NullTime{},
		},
	)
	transaction.Commit()
}

func UpdateTelegramContact(telegramId int, isSubscriber bool) {
	var contact telegramContact.TelegramContact
	contact.IsSubscriber = isSubscriber
	contact.TelegramId = telegramId
	transaction := DB.Connect().MustBegin()
	transaction.NamedExec(
		"UPDATE contact.telegram SET updated_at = :updated_at, is_subscriber = :is_subscriber "+
			" WHERE telegram_id = :telegram_id",
		&contact,
	)
	transaction.Commit()
}

func FindSubscribers() []telegramContact.TelegramContact {
	var contacts []telegramContact.TelegramContact
	query := "SELECT telegram_id FROM contact.telegram WHERE is_subscriber = true"
	err := DB.Connect().Select(&contacts, query)
	if err != nil {
		fmt.Println(err)
	}

	return contacts
}
