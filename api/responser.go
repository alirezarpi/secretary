package api

import (
	"encoding/json"
	"net/http"

	"github.com/google/uuid"
)

type jsonResponse struct {
	UUID   string                 `json:"uuid"`
	Sucess bool                   `json:"sucess"`
	Data   map[string]interface{} `json:"data"`
}

func Responser(w http.ResponseWriter, r *http.Request, status bool, statusCode int, response map[string]interface{}) {
	resp := jsonResponse{
		UUID:   uuid.New().String(),
		Sucess: status,
		Data:   response,
	}

	Middleware(w, r)
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(resp)
}
