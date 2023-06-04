package main

import (
    "fmt"
    "log"
    "net/http"
    "encoding/json"
    "github.com/gorilla/mux"
    "github.com/google/uuid"
)

type jsonResponse struct {
    id       string `json:"id"`
    status   string `json:"status"`
    data     string `json:"data"`
}

func init() {
    PORT := 8000
}

func handleRequests() {
    myRouter := mux.NewRouter().StrictSlash(true)
    myRouter.HandleFunc("/", homePage)
    myRouter.HandleFunc("/all", returnAllresponses)
    log.Fatal(http.ListenAndServe(":" + HTTP_PORT, myRouter))
}

func main() {
    fmt.Println("Rest API v2.0 - Mux Routers")
    responses = {
        jsonResponse{id: uuid.New(), data: "fuck off", status: "OK"},
    }
    handleRequests()
}

