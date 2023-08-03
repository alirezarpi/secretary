package api

import (
	"fmt"
	"log"
	"net/http"

	"secretary/alpha/internal"
	"secretary/alpha/utils"
)

func UserAPI(w http.ResponseWriter, r *http.Request) {
	if Middleware(w, r, false) {
		user := &internal.User{}
		switch r.Method {
		case "POST":
			reqBody, err := utils.HandleReqJson(r)
			if err != nil {
				log.Println(err)
				Responser(w, r, false, 400, map[string]interface{}{
					"error": "invalid data",
				})
				return
			}
			err = user.CreateUser(reqBody["username"].(string), reqBody["password"].(string), reqBody["active"].(bool))
			if err != nil {
				http.Error(w, fmt.Sprintf("Error creating user: %v", err), http.StatusInternalServerError)
				Responser(w, r, false, 400, map[string]interface{}{
					"error": "invalid data",
				})
				return
			}
			Responser(w, r, true, 201, map[string]interface{}{
				"user_data": "username " + reqBody["username"].(string) + " created successfully",
			})
			return
		case "GET":
			queryParam := r.URL.Query().Get("username")
			if queryParam == "" {
				Responser(w, r, true, 200, map[string]interface{}{
					"user_data": user.GetAllUsers(),
				})
				return
			} else {
				Responser(w, r, true, 200, map[string]interface{}{
					"user_data": user.GetUser(queryParam),
				})
				return
			}
		default:
			Responser(w, r, false, 405, map[string]interface{}{
				"error": "method not allowed",
			})
			return
		}
	} else {
		Responser(w, r, false, 401, map[string]interface{}{
			"error": "unauthorized",
		})
		return
	}
}
