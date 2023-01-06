package telegramContact

import (
	"database/sql"
)

type TelegramContacts struct {
	TelegramContacts []TelegramContact `json:"telegram_contact"`
}

type TelegramContact struct {
	TelegramId   int          `db:"telegram_id" json:"telegram_id" example:"42" format:"int"`
	TelegramName string       `db:"telegram_name" json:"telegram_name" example:"John Doe" format:"string"`
	IsSubscriber bool         `db:"is_subscriber" json:"is_subscriber" example:"false" format:"bool"`
	CreatedAt    sql.NullTime `db:"created_at" json:"created_at" example:"2022-01-01 00:00:00" format:"string" swaggertype:"string"`
	UpdatedAt    sql.NullTime `db:"updated_at" json:"updated_at" example:"2022-01-01 00:00:00" format:"string" swaggertype:"string"`
}
