package utils

import (
	"log"
	"runtime"
)

func Logger(logLevel string, message string) {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	_, file, line, ok := runtime.Caller(1)
	if !ok {
		file = "unknown"
		line = 0
	}
	switch logLevel {
	case "info":
		log.Printf("[INFO] %s:%d - %s\n", file, line, message)
	case "err":
		log.Printf("[ERROR] %s:%d - %s\n", file, line, message)
	}
}

