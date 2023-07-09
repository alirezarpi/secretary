package storage

import (
	"golang.org/x/crypto/bcrypt"
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
