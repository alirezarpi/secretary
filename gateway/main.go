package main

import (
	"flag"
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
	
	var listenAddr = flag.String("listenaddr", "0.0.0.0:6080", "secretary server address")
	flag.Parse()
	var handler = http.NewServeMux()

	handler.HandleFunc("/", home)
	handler.HandleFunc("/hz", api.HealthCheck)

	var server = &http.Server{
		Addr:    *listenAddr,
		Handler: handler,
	}

	log.Println("Starting server on", *listenAddr)
	log.Fatal(server.ListenAndServe())
}
