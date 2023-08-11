package main

import (
	"flag"
	"log"
	"net/http"

	"secretary/alpha/api"
	"secretary/alpha/internal"
	"secretary/alpha/storage"
)

func main() {
	var listenAddr = flag.String("listenaddr", "0.0.0.0:6080", "secretary server address")
	flag.Parse()

	storage.DatabaseInit()
	internal.RunFixtures()

	internal.ShowBanner("./banner.txt")

	var handler = http.NewServeMux()

	handler.HandleFunc("/", api.HomeAPI)

	handler.HandleFunc("/hz", api.HealthCheckAPI)

	handler.HandleFunc("/asksfor", api.AskAPI)

	handler.HandleFunc("/resource", api.ResourceAPI)

	handler.HandleFunc("/user", api.UserAPI)
	handler.HandleFunc("/user/self", api.SelfAPI)
	handler.HandleFunc("/user/login", api.LoginAPI)

	var server = &http.Server{
		Addr:    *listenAddr,
		Handler: handler,
	}

	log.Println("Starting server on", *listenAddr)
	log.Fatal(server.ListenAndServe())
}
