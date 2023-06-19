package utils

import (
	"time"
)

func CurrentTime() string {
	currentTime := time.Now().Format("2006-01-02 15:04:05")

	return currentTime
}
