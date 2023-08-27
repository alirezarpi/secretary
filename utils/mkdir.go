package utils

import (
	"os"
	"path/filepath"
)

func MakeDir(dir string) {
	newpath := filepath.Join(".", dir)
	err := os.MkdirAll(newpath, os.ModePerm)
	if err != nil {
		Logger("fatal", err.Error())
	}
}
