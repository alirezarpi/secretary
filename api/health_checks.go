package api

import (
	"net/http"
	"secretary/alpha/storage"
)

func HealthCheckAPI(w http.ResponseWriter, r *http.Request) {
	Middleware(w, r)
	Responser(w, r, true, 200, map[string]interface{}{
		"backend": map[string]interface{}{
			"success": true,
		},
		"database": map[string]interface{}{
			"success": storage.DatabaseHealthCheck(),
		},
	})
	return
}
