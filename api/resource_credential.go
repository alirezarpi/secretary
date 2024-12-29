package api

import (
	"net/http"

	"secretary/alpha/internal"
	"secretary/alpha/utils"
)

func ResourceCredentialAPI(w http.ResponseWriter, r *http.Request) {
	if Middleware(w, r) {
		resource_credential := &internal.ResourceCredential{}
		switch r.Method {
		case "POST":
			reqBody, err := utils.HandleReqJson(r)
			if err != nil {
				Responser(w, r, false, 400, map[string]interface{}{
					"error": "invalid data",
				})
				return
			}
			err, ru_uuid := resource_credential.CreateResourceCredential(
				reqBody["credential_id"].(string), 
				reqBody["resource_id"].(string),
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
				"resource_credential_data": "resource_credential " + ru_uuid + " created successfully",
			})
			return
		case "GET":
			queryCredID := r.URL.Query().Get("credential_id")
			queryResourceID := r.URL.Query().Get("resource_id")

			if queryCredID == "" && queryResourceID == "" {
				Responser(w, r, true, 200, map[string]interface{}{
					"resource_credential_data": resource_credential.GetAllResourceCredentials(),
				})
				return
			}
			if queryCredID != "" && queryResourceID != "" {
				Responser(w, r, true, 200, map[string]interface{}{
					"resource_credential_data": resource_credential.GetResourceCredential(queryCredID, queryResourceID),
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
