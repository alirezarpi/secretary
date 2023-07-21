package api

import (
	"log"
	"net/http"

	"secretary/alpha/internal"
	"secretary/alpha/utils"
)

func AskAPI(w http.ResponseWriter, r *http.Request) {
	if Middleware(w, r) {
		switch r.Method {
		case "POST":
			reqBody, err := utils.HandleReqJson(r)
			if err != nil {
				log.Println(err)
			}
			// FIXME Validators needed
			Responser(w, r, true, 200, map[string]interface{}{
				"ask_data": internal.CreateAsk(reqBody["what"].(string), reqBody["reason"].(string)),
			})
		case "GET":
			queryParam := r.URL.Query().Get("uuid")
			if queryParam == "" {
				Responser(w, r, true, 200, map[string]interface{}{
					"ask_data": internal.GetAsk(),
				})
			} else {
				Responser(w, r, true, 200, map[string]interface{}{
					"ask_data": internal.GetAsk(queryParam),
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
