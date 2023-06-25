package storage

import (
	"log"
	"fmt"
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
	
	initTables := []string{ // put columns too
		"asks_for", 
		"users",
	}

	for _, table := range initTables {
		query := fmt.Sprintf(`CREATE TABLE IF NOT EXISTS %s (
		  uuid TEXT NOT NULL PRIMARY KEY,
		  what TEXT NOT NULL,
		  created_time DATETIME NOT NULL,
		  modified_time DATETIME NOT NULL,
		  reason TEXT NOT NULL,
		  status TEXT NOT NULL
		);`, table)

		_, err := db.Exec(query)
		if err != nil {
			log.Fatal("Error in DatabaseInit table ", table, " : ", err)
			return false
		}
	}
	
	return true
}
