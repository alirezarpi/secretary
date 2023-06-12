package api

import (
	"net/http"
)


func Middleware(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

}

