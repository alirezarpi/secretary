package api

import (
	"fmt"
	"net/http"

	"secretary/alpha/internal/resource"
	"secretary/alpha/utils"
)

// curl -v -X POST -b $TOKEN  -H "Content-Type: application/json" -d '{"name": "test_db", "active": true, "dbType": "postgresql", "dbNames":"[*]", "dbPort": 5432, "dbHost": "localhost", "dbUser":"postgres", "dbPassword": "postgres"}' http://0.0.0.0:6080/db/resource | jq
func DatabaseResourceAPI(w http.ResponseWriter, r *http.Request) {
	if Middleware(w, r) {
		resource := &internal.DatabaseResource{}
		switch r.Method {
		case "POST":
			reqBody, err := utils.HandleReqJson(r)
			fmt.Println(reqBody["name"].(string),
				reqBody["active"].(bool),
				reqBody["dbType"].(string),
				reqBody["dbNames"].([]string),
				reqBody["dbPort"].(int),
				reqBody["dbHost"].(string),
				reqBody["dbUser"].(string),
				reqBody["dbPassword"].(string))

			if err != nil {
				Responser(w, r, false, 400, map[string]interface{}{
					"error": "invalid data",
				})
				return
			}
			err = resource.CreateDatabaseResource(
				reqBody["name"].(string),
				reqBody["active"].(bool),
				reqBody["dbType"].(string),
				reqBody["dbNames"].([]string),
				reqBody["dbPort"].(int),
				reqBody["dbHost"].(string),
				reqBody["dbUser"].(string),
				reqBody["dbPassword"].(string),
			)
			if err != nil {
				utils.Logger("err", err.Error())
				Responser(w, r, false, 400, map[string]interface{}{
					"error": err.Error(),
				})
				return
			}
			Responser(w, r, true, 201, map[string]interface{}{
				"resource_data": "name " + reqBody["name"].(string) + " created successfully",
			})
			return
		case "GET":
			queryParam := r.URL.Query().Get("name")
			if queryParam == "" {
				Responser(w, r, true, 200, map[string]interface{}{
					"resource_data": resource.GetAllDatabaseResources(),
				})
				return
			} else {
				Responser(w, r, true, 200, map[string]interface{}{
					"resource_data": resource.GetDatabaseResource(queryParam),
				})
				return
			}
		default:
			Responser(w, r, false, 405, map[string]interface{}{
				"error": "method not allowed",
			})
			return
		}
	}
}
