package internal

import (
	"fmt"
	"log"

	"golang.org/x/crypto/bcrypt"

	"secretary/alpha/storage"
	"secretary/alpha/utils"
)

type User struct {
	UUID         string
	Username     string
	PasswordHash string
	Active       bool
	CreatedTime  string
	ModifiedTime string
}

type SecureUser struct {
	UUID         string
	Username     string
	Active       bool
	CreatedTime  string
	ModifiedTime string
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

func (u *User) CreateUser(username string, password string, active bool) SecureUser {
	// NOTE add check on duplicate user
	// FIXME Validate
	createdTime := utils.CurrentTime()

	u.UUID = utils.UUID()
	u.Username = username
	u.CreatedTime = createdTime
	u.ModifiedTime = createdTime

	err := user.SetPassword(password)
	if err != nil {
		log.Fatal("SetPassword Error: ", err)
		return false
	}

	query := `
		INSERT INTO local_user (uuid, username, password_hash, active, created_time, updated_time)
		VALUES (?, ?, ?, ?, ?, ?)
	`
	_, err = storage.DatabaseExec(query, u.UUID, u.Username, u.PasswordHash, u.Active, u.CreatedTime, u.ModifiedTime)
	if err != nil {
		log.Fatal(err)
		return false
	}

	return User{
		UUID = u.UUID,
		Username = u.Username,
	}
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
