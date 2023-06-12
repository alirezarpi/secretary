package api

import (
	"log"
	"net/http"

	"secretary/alpha/utils"
)


func Ask(w http.ResponseWriter, r *http.Request) {
	Middleware(w, r)
	switch r.Method {
    case "POST":
		reqBody, err := utils.HandleReqJson(r)
		if err != nil {
			log.Fatal(err)
		} else {
			Responser(w, r, true, 200, map[string]interface{}{
				"echo": reqBody,
			})
		}
	default:
		Responser(w, r, true, 200, map[string]interface{}{
			"message": "TBD",
			},
		)
	}
}
