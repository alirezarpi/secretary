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

	handler.HandleFunc("/api/hz", api.HealthCheckAPI)

	handler.HandleFunc("/api/asksfor", api.AskAPI)

	handler.HandleFunc("/api/resource", api.ResourceAPI)

	handler.HandleFunc("/api/user", api.UserAPI)
	handler.HandleFunc("/api/user/self", api.SelfAPI)
	handler.HandleFunc("/api/user/login", api.LoginAPI)
	handler.HandleFunc("/api/user/logout", api.LogoutAPI)

	var server = &http.Server{
		Addr:    *listenAddr,
		Handler: handler,
	}

	log.Println("Server running on", *listenAddr)
	log.Fatal(server.ListenAndServe())
}
