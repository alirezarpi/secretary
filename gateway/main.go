package main

import (
	"log"
	"net/http"

	"github.com/alirezarpi/secretary/api"
)

type requestData struct {
	Inputs []map[string]interface{} `json:"data"`
}

func home(w http.ResponseWriter, r *http.Request) {
	api.responser(w, r, true, 200, map[string]interface{}{
		"message": "Secretary is here ^^",
	})
}

func main() {
	address := ":6080"
	handler := http.NewServeMux()

	handler.HandleFunc("/", home)
	handler.HandleFunc("/hz", api.healthCheck)

	server := &http.Server{
		Addr:    address,
		Handler: handler,
	}

	log.Println("Starting server on", address)
	log.Fatal(server.ListenAndServe())
}
