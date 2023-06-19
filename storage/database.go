package storage

import (
	"log"
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

func OpenDatabase() (*sql.DB) {
	db, err := sql.Open("sqlite3", "./storage/secretary.db")
	if err != nil {
		log.Fatal(err)
		return nil
	}
	return db
}

func DatabaseInit() bool {
	db := OpenDatabase()

	query := `CREATE TABLE IF NOT EXISTS asks_for (
	  uuid TEXT NOT NULL PRIMARY KEY,
	  what TEXT NOT NULL,
	  created_time DATETIME NOT NULL,
	  modified_time DATETIME NOT NULL,
	  reason TEXT NOT NULL,
	  status TEXT NOT NULL
	);`

	_, err := db.Exec(query)
	if err != nil {
		log.Fatal("Error in DatabaseInit: ", err)
		return false
	}

	return true
}
