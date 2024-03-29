package internal

import (
	"io/ioutil"
	"log"
)

func ShowBanner(filepath string) {
	content, err := ioutil.ReadFile("banner.txt")
	if err != nil {
		log.Println("Error reading the banner.txt file:", err)
	}
	log.Println("\n" + string(content))
}
