package db

import (
	"log"
	"os"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func InitDb() *sqlx.DB {
	db, err := sqlx.Connect("postgres", os.Getenv("DB_DBSTRING"))
	if err != nil {
		// you should not be able to run the app if you cannot connect to the database
		log.Fatalf("failed to connect to the database: %v", err)
	}
	return db
}
