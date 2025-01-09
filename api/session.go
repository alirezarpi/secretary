package api

import (
	"net/http"

	"secretary/alpha/internal"
	"secretary/alpha/utils"
)

func SessionAPI(w http.ResponseWriter, r *http.Request) {
	if Middleware(w, r) {
		session := &internal.Session{}
		switch r.Method {
		case "POST":
			reqBody, err := utils.HandleReqJson(r)
			if err != nil {
				Responser(w, r, false, 400, map[string]interface{}{
					"error": "invalid data",
				})
				return
			}
			err, session_uuid := session.CreateSession(
				reqBody["name"].(string),
				reqBody["host"].(string),
				reqBody["port"].(string),
				reqBody["kind"].(string),
				reqBody["active"].(bool),
			)
			if err != nil {
				utils.Logger("err", err.Error())
				Responser(w, r, false, 400, map[string]interface{}{
					"error": err.Error(),
				})
				return
			}
			Responser(w, r, true, 201, map[string]interface{}{
				"session_data": "session " + session_uuid + " created successfully",
			})
			return
		case "GET":
			queryParam := r.URL.Query().Get("name")
			if queryParam == "" {
				Responser(w, r, true, 200, map[string]interface{}{
					"session_data": session.GetAllSessions(),
				})
				return
			} else {
				Responser(w, r, true, 200, map[string]interface{}{
					"session_data": session.GetSession(queryParam),
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
