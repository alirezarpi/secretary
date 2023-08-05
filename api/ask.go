package api

import (
	"net/http"

	"secretary/alpha/internal"
	"secretary/alpha/utils"
)

func AskAPI(w http.ResponseWriter, r *http.Request) {
	if Middleware(w, r) {
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
			if (err != nil) {
				println("fdfdsfs")
				Responser(w, r, false, 500, map[string]interface{}{
					"error": "internal error",
				})
				return
			}
			Responser(w, r, true, 201, map[string]interface{}{
				"ask_data": "asksfor created successfully",
			})
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
