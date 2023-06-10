package api

import (
	"net/http"
)


func HealthCheck(w http.ResponseWriter, r *http.Request) {
	Responser(w, r, true, 200, map[string]interface{}{
		"backend": map[string]interface{}{
			"message": "healthy",
			"success": true,
		},
	})
}


