package main

import (
	"flag"
	"log"
	"net/http"

	"secretary/alpha/api"
	"secretary/alpha/storage"
)


func main() {
	var listenAddr = flag.String("listenaddr", "0.0.0.0:6080", "secretary server address")
	flag.Parse()
	storage.DatabaseInit()
	var handler = http.NewServeMux()

	handler.HandleFunc("/", api.Home)
	handler.HandleFunc("/hz", api.HealthCheck)
	handler.HandleFunc("/ask", api.Ask)

	var server = &http.Server{
		Addr:    *listenAddr,
		Handler: handler,
	}

	log.Println("Starting server on", *listenAddr)
	log.Fatal(server.ListenAndServe())
}
