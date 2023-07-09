package internal

import (
	"fmt"
	"log"

	"secretary/alpha/storage"
	"secretary/alpha/utils"
)


func CreateUser(username string, password string, active bool) bool {
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

	// FIXME https://stackoverflow.com/a/36164458
	query := fmt.Sprintf(`
		INSERT INTO local_user (uuid, username, password_hash, active, created_date, modified_time)
		VALUES (%s, %s, %s, %s, %s, %s)
	`, user.UUID, user.Username, user.PasswordHash, user.Active, user.CreatedTime, user.ModifiedTime)

	_, err = storage.DatabaseExec(query)
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

