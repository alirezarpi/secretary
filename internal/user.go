package internal

import (
	"fmt"
	"log"

	"secretary/alpha/storage"
	"secretary/alpha/utils"
)


func CreateUser(username string, password string, active bool) bool {
	// NOTE add check on duplicate user
	// FIXME Validate
	uuid := utils.UUID()
	createdTime := utils.CurrentTime()

	user := storage.User{
		UUID: uuid,
		Username: username,
		Active: active,
		CreatedTime: createdTime,
		ModifiedTime: createdTime,
	}

	err := user.SetPassword(password)
	if err != nil {
		log.Fatal("SetPassword Error: ", err)
		return false
	}

	query := `
		INSERT INTO local_user (uuid, username, password_hash, active, created_time, updated_time)
		VALUES (?, ?, ?, ?, ?, ?)
	`
	_, err = storage.DatabaseExec(query, user.UUID, user.Username, user.PasswordHash, user.Active, user.CreatedTime, user.ModifiedTime)
	if err != nil {
		log.Fatal(err)
		return false
	}

	return true
}

func GetUser(username ...string) []map[string]interface{} {
	var query string
	if len(username) > 0 {
		query = fmt.Sprintf(`SELECT * FROM local_user WHERE username='%s'`, username[0])
	} else {
		query = `SELECT * FROM local_user`
	}

	rows, err := storage.DatabaseQuery(query)
	if err != nil {
		log.Fatal("Error in GetUser: ", err)
		return []map[string]interface{}{}
	}
	defer rows.Close()

	columns, err := rows.Columns()
	if err != nil {
		log.Fatal("Error in GetUser: ", err)
		return []map[string]interface{}{}
	}

	results, err := utils.HandleTableToJSON(columns, rows)
	if err != nil {
		log.Fatal("Error in GetUser: ", err)
		return []map[string]interface{}{}
	}

	return results
}

