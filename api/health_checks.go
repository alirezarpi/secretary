



func healthCheck(w http.ResponseWriter, r *http.Request) {
	responser(w, r, true, 200, map[string]interface{}{
		"backend": map[string]interface{}{
			"message": "healthy",
			"success": true,
		},
	})
}


