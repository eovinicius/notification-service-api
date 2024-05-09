package config

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

func DbContext() *sql.DB {
	db, err := sql.Open("sqlite3", "./database.db")
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	return db
}
