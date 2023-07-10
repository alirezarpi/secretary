package main

import (
	"flag"
	"log"
	"net/http"

	"secretary/alpha/api"
	"secretary/alpha/storage"
	"secretary/alpha/internal"
)


func main() {
	var listenAddr = flag.String("listenaddr", "0.0.0.0:6080", "secretary server address")
	flag.Parse()
	storage.DatabaseInit()
	internal.ShowBanner("./banner.txt")
	
	var handler = http.NewServeMux()

	handler.HandleFunc("/", api.HomeAPI)
	handler.HandleFunc("/hz", api.HealthCheckAPI)
	handler.HandleFunc("/ask", api.AskAPI)
	handler.HandleFunc("/user", api.UserAPI)

	var server = &http.Server{
		Addr:    *listenAddr,
		Handler: handler,
	}
	
	log.Println("Starting server on", *listenAddr)
	log.Fatal(server.ListenAndServe())
}
