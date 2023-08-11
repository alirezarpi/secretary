package utils

import (
	"math/rand"
	"time"
)

const (
	letterBytes  = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	digitBytes   = "0123456789"
	specialBytes = "!@#$%^&*()-_=+[]{}|;:,.<>?~"
)

func GenerateRandomPassword(length int) string {
	rand.Seed(time.Now().UnixNano())

	chars := letterBytes + digitBytes + specialBytes
	password := make([]byte, length)
	for i := range password {
		password[i] = chars[rand.Intn(len(chars))]
	}
	return string(password)
}
