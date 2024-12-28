package api

import (
	"net/http"

	"secretary/alpha/internal"
	"secretary/alpha/utils"
)

func ResourceUserAPI(w http.ResponseWriter, r *http.Request) {
	if Middleware(w, r) {
		resource_user := &internal.ResourceUser{}
		switch r.Method {
		case "POST":
			reqBody, err := utils.HandleReqJson(r)
			if err != nil {
				Responser(w, r, false, 400, map[string]interface{}{
					"error": "invalid data",
				})
				return
			}
			err, ru_uuid := resource_user.CreateResourceUser(
				reqBody["user_id"].(string), 
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
				"resource_user_data": "resource_user " + ru_uuid + " created successfully",
			})
			return
		case "GET":
			queryUserID := r.URL.Query().Get("user_id")
			queryResourceID := r.URL.Query().Get("resource_id")

			if queryUserID == "" && queryResourceID == "" {
				Responser(w, r, true, 200, map[string]interface{}{
					"resource_user_data": resource_user.GetAllResourceUsers(),
				})
				return
			}
			if queryUserID != "" && queryResourceID != "" {
				Responser(w, r, true, 200, map[string]interface{}{
					"resource_user_data": resource_user.GetResourceUser(queryUserID, queryResourceID),
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
