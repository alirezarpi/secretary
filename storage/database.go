package storage

import (
	"database/sql"
	"fmt"

	"secretary/alpha/utils"
)

func OpenDatabase() *sql.DB {
	utils.MakeDir("persistence")
	db, err := sql.Open("sqlite3", "./storage/secretary.db")
	if err != nil {
		utils.Logger("fatal", err.Error())
		return nil
	}
	return db
}

func DatabaseInit() bool {
	db := OpenDatabase()

	// ASKSFOR Tables
	table := "asksfor"
	query := fmt.Sprintf(`
		CREATE TABLE IF NOT EXISTS %s (
			uuid TEXT NOT NULL PRIMARY KEY,
			what TEXT NOT NULL,
			reason TEXT NOT NULL,
			status TEXT NOT NULL,
			requester TEXT NOT NULL,
			reviewer TEXT NOT NULL,
			created_time TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
			modified_time TIMESTAMP DEFAULT CURRENT_TIMESTAMP
		);`, table)
	_, err := db.Exec(query)
	if err != nil {
		utils.Logger("fatal", err.Error())
		return false
	}

	// USER Tables
	table = "user_local"
	query = fmt.Sprintf(`
		CREATE TABLE IF NOT EXISTS %s (
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

	// Resource Tables
	table = "resource"
	query = fmt.Sprintf(`
		CREATE TABLE IF NOT EXISTS %s (
			uuid TEXT NOT NULL PRIMARY KEY,
			name TEXT NOT NULL,
			active BOOLEAN DEFAULT TRUE,
			created_time TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
			modified_time TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
			UNIQUE (uuid, name)
		);`, table)
	_, err = db.Exec(query)
	if err != nil {
		utils.Logger("fatal", err.Error())
		return false
	}

	// RBAC Tables
	table = "rbac_role"
	query = fmt.Sprintf(`
		CREATE TABLE IF NOT EXISTS %s (
			uuid TEXT NOT NULL PRIMARY KEY,
			name TEXT NOT NULL,
			active BOOLEAN DEFAULT TRUE,
			created_time TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
			modified_time TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
			UNIQUE (uuid, name)
		);`, table)
	_, err = db.Exec(query)
	if err != nil {
		utils.Logger("fatal", err.Error())
		return false
	}

	table = "rbac_permissions"
	query = fmt.Sprintf(`
		CREATE TABLE IF NOT EXISTS %s (
			uuid TEXT NOT NULL PRIMARY KEY,
			name TEXT NOT NULL,
			active BOOLEAN DEFAULT TRUE,
			created_time TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
			modified_time TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
			UNIQUE (uuid, name)
		);`, table)
	_, err = db.Exec(query)
	if err != nil {
		utils.Logger("fatal", err.Error())
		return false
	}

	table = "rbac_user_role"
	query = fmt.Sprintf(`
		CREATE TABLE IF NOT EXISTS %s (
			uuid TEXT NOT NULL PRIMARY KEY,
			user_username TEXT NOT NULL,
			role_name TEXT NOT NULL,
			active BOOLEAN DEFAULT TRUE,
			created_time TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
			modified_time TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
			UNIQUE (uuid)
		);`, table)
	_, err = db.Exec(query)
	if err != nil {
		utils.Logger("fatal", err.Error())
		return false
	}

	utils.Logger("info", "Database successfully initiated")
	return true
}
