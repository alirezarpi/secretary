package api

import (
	"net/http"
)

func HealthCheckAPI(w http.ResponseWriter, r *http.Request) {
	if (Middleware(w, r)) {
		Responser(w, r, true, 200, map[string]interface{}{
			"backend": map[string]interface{}{
				"success": true,
			},
			"database": map[string]interface{}{
				"success": true,//storage.DatabaseHealthCheck(),
			},
		})
		return
	} else {
		Responser(w, r, true, 200, map[string]interface{}{
			"message": "unauthorized",
		})
		return
	}
}
