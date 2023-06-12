package storage

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"log"
)

func DatabaseInit() {
	db, err := sql.Open("sqlite3", "./storage/secretary.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	sqlStmt := `CREATE TABLE IF NOT EXISTS asks_for (
	  uuid TEXT NOT NULL PRIMARY KEY,
	  what TEXT NOT NULL,
	  created_time DATETIME NOT NULL,
	  modified_time DATETIME NOT NULL,
	  reason TEXT NOT NULL,
	  status TEXT NOT NULL
	);`

	_, err = db.Exec(sqlStmt)
	if err != nil {
		log.Printf("%q: %s\n", err, sqlStmt)
		return
	}
}
