package api

import (
	"net/http"

	"secretary/alpha/internal"
	"secretary/alpha/utils"
)

func ResourceDatabaseAPI(w http.ResponseWriter, r *http.Request) {
	if Middleware(w, r) {
		resource_database := &internal.ResourceDatabase{}
		switch r.Method {
		case "POST":
			reqBody, err := utils.HandleReqJson(r)
			if err != nil {
				Responser(w, r, false, 400, map[string]interface{}{
					"error": "invalid data",
				})
				return
			}
			err, rd_uuid := resource_database.CreateResourceDatabase(
				reqBody["name"].(string),
				reqBody["resource_user_id"].(string),
				reqBody["db_host"].(string), 
				reqBody["db_port"].(string), 
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
				"resource_database_data": "resource_database " + rd_uuid + " created successfully",
			})
			return
		case "GET":
			println("FIXME HERE")
			//queryUUID := r.URL.Query().Get("uuid")

			//FIXME Fix this
			//if queryUserID == "" {
			//	Responser(w, r, true, 200, map[string]interface{}{
			//		"resource_database_data": resource_user.GetAllResourceUsers(),
			//	})
			//	return
			//} else if queryUUID != "" {
			//	Responser(w, r, true, 200, map[string]interface{}{
			//		"resource_database_data": resource_database.GetResourceDatabase(queryUUID),
			//	})
			//	return
			//} else {
			//	Responser(w, r, false, 400, map[string]interface{}{
			//		"error": "invalid input -- queryparam",
			//	})
			//	return
			//}
		default:
			Responser(w, r, false, 405, map[string]interface{}{
				"error": "method not allowed",
			})
			return
		}
	}
}
