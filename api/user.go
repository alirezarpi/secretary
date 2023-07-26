package api

import (
	"fmt"
	"log"
	"net/http"

	"secretary/alpha/internal"
	"secretary/alpha/utils"
)

func UserAPI(w http.ResponseWriter, r *http.Request) {
	if Middleware(w, r) {
		user := &internal.User{}
		switch r.Method {
		case "POST":
			reqBody, err := utils.HandleReqJson(r)
			if err != nil {
				log.Println(err)
				Responser(w, r, false, 400, map[string]interface{}{
					"error": "invalid data",
				})
			}
			err = user.CreateUser(reqBody["username"].(string), reqBody["password"].(string), reqBody["active"].(bool))
			if err != nil {
				http.Error(w, fmt.Sprintf("Error creating user: %v", err), http.StatusInternalServerError)
				Responser(w, r, false, 400, map[string]interface{}{
					"error": "invalid data",
				})
			}
			Responser(w, r, true, 201, map[string]interface{}{
				"user_data": "username " + reqBody["username"].(string) + " created successfully",
			})
		case "GET":
			queryParam := r.URL.Query().Get("username")
			if queryParam == "" {
				Responser(w, r, true, 200, map[string]interface{}{
					"user_data": user.GetAllUsers(),
				})
			} else {
				Responser(w, r, true, 200, map[string]interface{}{
					"user_data": user.GetUser(queryParam),
				})
			}
		default:
			Responser(w, r, false, 405, map[string]interface{}{
				"error": "method not allowed",
			})
		}
	} else {
		Responser(w, r, false, 401, map[string]interface{}{
			"error": "Unauthorized",
		})
	}
}
