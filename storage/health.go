package storage

import (
	"log"

	_ "github.com/mattn/go-sqlite3"
)

func DatabaseHealthCheck() bool {
	db := OpenDatabase()
	err := db.Ping()
	if err != nil {
		log.Fatalf("Database connection health check failed: %v", err)
		return false
	}
	return true
}
