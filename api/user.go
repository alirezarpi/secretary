package api

import (
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
				Responser(w, r, false, 400, map[string]interface{}{
					"error": "invalid data",
				})
				return
			}
			err = user.CreateUser(reqBody["username"].(string), reqBody["password"].(string), reqBody["active"].(bool))
			if err != nil {
				utils.Logger("err", err.Error())
				Responser(w, r, false, 400, map[string]interface{}{
					"error": err.Error(),
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
	}
}
