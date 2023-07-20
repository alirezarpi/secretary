package api

import (
	"log"
	"net/http"

	"secretary/alpha/internal"
	"secretary/alpha/utils"
)


func UserAPI(w http.ResponseWriter, r *http.Request) {
	if (Middleware(w, r)) {
		switch r.Method {
		case "POST":
			reqBody, err := utils.HandleReqJson(r)
			if err != nil {
				log.Println(err)
				Responser(w, r, true, 400, map[string]interface{}{
					"message": "invalid data",
				})
			}
			// FIXME Validators needed
			Responser(w, r, true, 201, map[string]interface{}{
				"user_data": internal.CreateUser(reqBody["username"].(string), reqBody["password"].(string), reqBody["active"].(bool)),
			})
		case "GET":
			queryParam := r.URL.Query().Get("username")
			if queryParam == "" {
				Responser(w, r, true, 200, map[string]interface{}{
					"user_data": internal.GetUser(),
				})
			} else {
				Responser(w, r, true, 200, map[string]interface{}{
					"user_data": internal.GetUser(queryParam),
				})
			}
		default:
			Responser(w, r, false, 405, map[string]interface{}{
				"message": "method not allowed",
			})
		}
	} else {
		Responser(w, r, false, 401, map[string]interface{}{
			"message": "Unauthorized",
		})
	}
}
