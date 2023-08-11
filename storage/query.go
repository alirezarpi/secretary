package storage

import (
	"database/sql"

	"secretary/alpha/utils"

	_ "github.com/mattn/go-sqlite3"
)

func DatabaseQuery(query string) (*sql.Rows, error) {
	db := OpenDatabase()

	rows, err := db.Query(query)
	if err != nil {
		utils.Logger("err", err.Error())
		return nil, err
	}
	defer db.Close()

	return rows, nil
}

func DatabaseExec(query string, args ...interface{}) (*sql.Result, error) {
	db := OpenDatabase()

	result, err := db.Exec(query, args...)
	if err != nil {
		return nil, err
	}
	defer db.Close()

	utils.Logger("info", "query done: "+query)
	return &result, nil
}
