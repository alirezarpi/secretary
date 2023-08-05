package storage

import (
	_ "github.com/mattn/go-sqlite3"

	"secretary/alpha/utils"
)

func DatabaseHealthCheck() bool {
	db := OpenDatabase()
	err := db.Ping()
	if err != nil {
		utils.Logger("fatal", err.Error())
		return false
	}
	return true
}
