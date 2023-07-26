package internal

import (
	"fmt"
	"log"
	"time"

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

func (u *User) CreateUser(username string, password string, active bool) error {
	existingUser := u.GetUser(username)
	if existingUser != nil {
		return fmt.Errorf("username already exists")
	}

	// FIXME Add validation code here ...
	// FIXME change the error handling

	createdTime := utils.CurrentTime()

	u.UUID = utils.UUID()
	u.Username = username
	u.CreatedTime = createdTime
	u.ModifiedTime = createdTime
	u.Active = active

	err := u.SetPassword(password)
	if err != nil {
		return fmt.Errorf("SetPassword Error: %v", err)
	}

	query := `
		INSERT INTO local_user (uuid, username, password_hash, active, created_time, modified_time)
		VALUES (?, ?, ?, ?, ?, ?)
	`
	_, err = storage.DatabaseExec(query, u.UUID, u.Username, u.PasswordHash, u.Active, u.CreatedTime, u.ModifiedTime)
	if err != nil {
		return fmt.Errorf("Error in CreateUser: %v", err)
	}

	return nil
}

func (u *User) GetUser(username string) *User {
	query := fmt.Sprintf(`SELECT * FROM local_user WHERE username='%s'`, username)

	rows, err := storage.DatabaseQuery(query)
	if err != nil {
		log.Fatal("Error in GetUser: ", err)
		return nil
	}
	defer rows.Close()

	columns, err := rows.Columns()
	if err != nil {
		log.Fatal("Error in GetUser: ", err)
		return nil
	}

	results, err := utils.HandleTableToJSON(columns, rows)
	if err != nil {
		log.Fatal("Error in GetUser: ", err)
		return nil
	}

	if len(results) == 0 {
		return nil
	}

	return &User{
		UUID:         results[0]["uuid"].(string),
		Username:     results[0]["username"].(string),
		PasswordHash: results[0]["password_hash"].(string),
		Active:       results[0]["active"].(bool),
		CreatedTime:  results[0]["created_time"].(time.Time).Format(time.RFC3339),
		ModifiedTime: results[0]["modified_time"].(time.Time).Format(time.RFC3339),
	}
}

func (u *User) GetAllUsers() []*User {
	query := `SELECT * FROM local_user`

	rows, err := storage.DatabaseQuery(query)
	if err != nil {
		log.Fatal("Error in GetUsers: ", err)
		return nil
	}
	defer rows.Close()

	columns, err := rows.Columns()
	if err != nil {
		log.Fatal("Error in GetUsers: ", err)
		return nil
	}

	results, err := utils.HandleTableToJSON(columns, rows)
	if err != nil {
		log.Fatal("Error in GetUsers: ", err)
		return nil
	}

	users := make([]*User, 0, len(results))
	for _, res := range results {
		println(res["created_time"].(time.Time).Format(time.RFC3339))
		println(res["modified_time"].(time.Time).Format(time.RFC3339))
		user := &User{
			UUID:         res["uuid"].(string),
			Username:     res["username"].(string),
			PasswordHash: res["password_hash"].(string),
			Active:       res["active"].(bool),
			CreatedTime:  res["created_time"].(time.Time).Format(time.RFC3339),
			ModifiedTime: res["modified_time"].(time.Time).Format(time.RFC3339),
		}
		users = append(users, user)
	}

	return users
}
