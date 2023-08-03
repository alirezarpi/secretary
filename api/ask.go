package api

import (
	"net/http"

	"secretary/alpha/utils"
	"secretary/alpha/internal"
)

func AskAPI(w http.ResponseWriter, r *http.Request) {
	if Middleware(w, r) {
		asksFor := &internal.AsksFor{}
		switch r.Method {
		case "POST":
			reqBody, err := utils.HandleReqJson(r)
			if err != nil {
				utils.ErrorLogger(err)
			}

			user := &internal.User{}
			_, user = isAuthenticated(r)
			// FIXME Validators needed, data and also check if reviewer is valid user
			Responser(w, r, true, 200, map[string]interface{}{
				"ask_data": asksFor.CreateAsksFor(
					reqBody["what"].(string),
					reqBody["reason"].(string),
					user.Username,
					reqBody["reviewer"].(string),
				)})
			return
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
	} else {
		Responser(w, r, false, 401, map[string]interface{}{
			"error": "unauthorized",
		})
		return
	}
}
