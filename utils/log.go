package utils

import (
	"log"
)

func Logger(logLevel string, message string) {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	switch logLevel {
	case "info":
		log.Println(message)
	}
}

func ErrorLogger(err error) {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Panic(err)
}

