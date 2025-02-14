package api

import (
	"net/http"

	"secretary/alpha/internal"
	"secretary/alpha/utils"
)

func ResourceAPI(w http.ResponseWriter, r *http.Request) {
	if Middleware(w, r) {
		resource := &internal.Resource{}
		switch r.Method {
		case "POST":
			reqBody, err := utils.HandleReqJson(r)
			if err != nil {
				Responser(w, r, false, 400, map[string]interface{}{
					"error": "invalid data",
				})
				return
			}
			err, resource_uuid := resource.CreateResource(
				reqBody["name"].(string),
				reqBody["host"].(string),
				reqBody["port"].(string),
				reqBody["kind"].(string),
				reqBody["active"].(bool),
			)
			if err != nil {
				utils.Logger("err", err.Error())
				Responser(w, r, false, 400, map[string]interface{}{
					"error": err.Error(),
				})
				return
			}
			Responser(w, r, true, 201, map[string]interface{}{
				"resource_data": "resource " + resource_uuid + " created successfully",
			})
			return
		case "GET":
			queryParam := r.URL.Query().Get("name")
			if queryParam == "" {
				Responser(w, r, true, 200, map[string]interface{}{
					"resource_data": resource.GetAllResources(),
				})
				return
			} else {
				Responser(w, r, true, 200, map[string]interface{}{
					"resource_data": resource.GetResource(queryParam),
				})
				return
			}
		default:
			Responser(w, r, false, 405, map[string]interface{}{
				"error": "method not allowed",
			})
			return
		}
	}
}
