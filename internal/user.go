package internal

import (
	"fmt"
	"log"

	"golang.org/x/crypto/bcrypt"

	"secretary/alpha/storage"
	"secretary/alpha/utils"
)

type User struct {
	UUID			string
	Username		string
	PasswordHash	string
	Active			bool
	CreatedTime		string
	ModifiedTime	string
}


func (u *User) SetPassword(password string) error {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	u.PasswordHash = string(hash)
	return nil
}

func (u *User) CheckPassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.PasswordHash), []byte(password))
	return err == nil
}

func CreateUser(username string, password string, active bool) bool {
	// NOTE add check on duplicate user
	// FIXME Validate
	uuid := utils.UUID()
	createdTime := utils.CurrentTime()

	user := User{
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

