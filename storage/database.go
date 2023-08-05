package storage

import (
	"database/sql"
	"fmt"
	"log"
	"secretary/alpha/utils"

	_ "github.com/mattn/go-sqlite3"
)

func OpenDatabase() *sql.DB {
	db, err := sql.Open("sqlite3", "./storage/secretary.db")
	if err != nil {
		utils.Logger("fatal", err.Error())
		return nil
	}
	return db
}

func DatabaseInit() bool {
	db := OpenDatabase()

	table := "asks_for"
	query := fmt.Sprintf(`
	CREATE TABLE IF NOT EXISTS %s (
		uuid TEXT NOT NULL PRIMARY KEY,
		what TEXT NOT NULL,
		created_time DATETIME NOT NULL,
		modified_time DATETIME NOT NULL,
		reason TEXT NOT NULL,
		status TEXT NOT NULL,
		requester TEXT NOT NULL,
		reviewer TEXT NOT NULL
	);`, table)

	_, err := db.Exec(query)
	if err != nil {
		utils.Logger("fatal", err.Error())
		return false
	}

	table = "local_user"
	query = fmt.Sprintf(`CREATE TABLE IF NOT EXISTS %s (
		uuid TEXT NOT NULL PRIMARY KEY,
		username TEXT NOT NULL,
		password_hash TEXT NOT NULL,
		active BOOLEAN DEFAULT TRUE,
		created_time TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
		modified_time TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
		UNIQUE (uuid, username)
	);`, table)

	_, err = db.Exec(query)
	if err != nil {
		utils.Logger("fatal", err.Error())
		return false
	}

	log.Println("Database successfully initiated")
	return true
}
