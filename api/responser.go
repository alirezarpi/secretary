package api

import (
	"encoding/json"
	"fmt"
	"reflect"
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

	w.WriteHeader(statusCode)
	fmt.Println(reflect.TypeOf(resp).Kind())
	fmt.Println(reflect.TypeOf(resp.Data["message"]).Kind())
	json.NewEncoder(w).Encode(resp)
}
