package database

import (
	"database/sql"
	"log"
)

func NewDatabaseConnection() *sql.DB {
	database, errorDatabase := sql.Open("sqlLite3", "./tasks.db")
	if errorDatabase != nil {
		log.Fatalf("Error to connect database: %v", errorDatabase)
	}

	return database
}
