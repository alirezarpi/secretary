package api

import (
	"net/http"
)

func Home(w http.ResponseWriter, r *http.Request) {
	Middleware(w, r)
	if r.URL.Path != "/" {
		Responser(w, r, false, 404, map[string]interface{}{
		"message": "path not found",
		})
		return
	}
	Responser(w, r, true, 200, map[string]interface{}{
		"message": "Secretary is here ^^",
	})
}
