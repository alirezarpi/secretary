package api

import (
	"net/http"

	"secretary/alpha/internal"
	"secretary/alpha/internal/audit"
	"secretary/alpha/utils"
)

var MiddlewareWrapper = Middleware

func AskAPI(w http.ResponseWriter, r *http.Request) {
	if MiddlewareWrapper(w, r) {
		asksFor := &internal.AsksFor{}
		switch r.Method {
		case "POST":
			reqBody, err := utils.HandleReqJson(r)
			if err != nil {
				utils.Logger("err", err.Error())
				Responser(w, r, false, 400, map[string]interface{}{
					"error": "invalid data",
				})
				return
			}

			user := &internal.User{}
			_, user = isAuthenticated(r)
			err = asksFor.CreateAsksFor(
				reqBody["what"].(string),
				reqBody["reason"].(string),
				user.Username,
				reqBody["reviewer"].(string),
			)
			if err != nil {
				Responser(w, r, false, 500, map[string]interface{}{
					"error": "internal error",
				})
				return
			}
			audit.Audit("[asksfor] [action:create] user " + user.Username + " logged in.")
			Responser(w, r, true, 201, map[string]interface{}{
				"ask_data": "asksfor created successfully",
			})
			return
		case "PATCH":
			utils.Logger("info", "Requesting for patch the asksfor")
		case "GET":
			queryParam := r.URL.Query().Get("uuid")
			if queryParam == "" {
				Responser(w, r, true, 200, map[string]interface{}{
					"ask_data": asksFor.GetAllAsksFors(),
				})
				return
			} else {
				Responser(w, r, true, 200, map[string]interface{}{
					"ask_data": asksFor.GetAsksFor(queryParam),
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
