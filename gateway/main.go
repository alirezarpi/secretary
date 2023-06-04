package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/google/uuid"
)

type jsonResponse struct {
	UUID   string       `json:"uuid"`
	Status string       `json:"status"`
	Data   []string     `json:"data"`
}

func homePage(w http.ResponseWriter, r *http.Request) {
    status := "OK"
	response := jsonResponse{
		UUID:   uuid.New().String(),
		Status: status,
        Data: []string{{"fdsf":"fsdf"}},
    }

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func main() {
    address := ":6080"
    handler := http.NewServeMux()

    handler.HandleFunc("/", homePage)

    server := &http.Server{
        Addr:    address,
        Handler: handler,
    }

    log.Println("Starting server on", address)
    log.Fatal(server.ListenAndServe())
}
