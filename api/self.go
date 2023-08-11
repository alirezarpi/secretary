package api

import (
	"net/http"

	"secretary/alpha/internal"
)

func SelfAPI(w http.ResponseWriter, r *http.Request) {
	if Middleware(w, r) {
		if r.Method != "GET" {
			Responser(w, r, false, 405, map[string]interface{}{
				"error": "method not allowed",
			})
			return
		}

		user := &internal.User{}
		_, user = isAuthenticated(r)
		Responser(w, r, true, 200, map[string]interface{}{
			"user_data": user,
		})
		return
	} else {
		Responser(w, r, false, 401, map[string]interface{}{
			"error": "unauthorized",
		})
		return
	}
}
