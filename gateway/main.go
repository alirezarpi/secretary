package main

import (
	"log"
	"net/http"
	
	"secretary/alpha/api"
)

type requestData struct {
	Inputs []map[string]interface{} `json:"data"`
}

func home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		api.Responser(w, r, false, 404, map[string]interface{}{
		"message": "path not found",
		})
		return
	}
	api.Responser(w, r, true, 200, map[string]interface{}{
		"message": "Secretary is here ^^",
	})
}

func main() {
	address := ":6080"
	handler := http.NewServeMux()

	handler.HandleFunc("/", home)
	handler.HandleFunc("/hz", api.HealthCheck)

	server := &http.Server{
		Addr:    address,
		Handler: handler,
	}

	log.Println("Starting server on", address)
	log.Fatal(server.ListenAndServe())
}
