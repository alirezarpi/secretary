package api

import (
	"net/http"
)

func HomeAPI(w http.ResponseWriter, r *http.Request) {
	Middleware(w, r, false)
	if r.Method != "GET" {
		Responser(w, r, false, 405, map[string]interface{}{
			"message": "method not allowed",
		})
		return
	}
	if r.URL.Path != "/" {
		Responser(w, r, false, 404, map[string]interface{}{
			"message": "path not found",
		})
		return
	}
	Responser(w, r, true, 200, map[string]interface{}{
		"message": "Secretary is here ^^",
	})
	return
}
