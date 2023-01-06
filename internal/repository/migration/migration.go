package migration

import (
	_ "database/sql"
	_ "github.com/lib/pq"
	"golang_telegram_bot/internal/database/DB"
	"log"
)

type Migration struct {
	Id        int    `db:"id"`
	Timestamp int    `db:"timestamp"`
	Name      string `db:"name"`
}

func FindAllMigration() []Migration {
	var migration []Migration
	err := DB.Connect().Select(&migration, "SELECT * FROM migrations")
	if err != nil {
		log.Fatalln(err)
	}
	return migration
}
