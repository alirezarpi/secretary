package api

import (
	"net/http"

	"secretary/alpha/internal"
	"secretary/alpha/utils"
)

func CredentialAPI(w http.ResponseWriter, r *http.Request) {
	if Middleware(w, r) {
		credential := &internal.Credential{}
		switch r.Method {
		case "POST":
			reqBody, err := utils.HandleReqJson(r)
			if err != nil {
				Responser(w, r, false, 400, map[string]interface{}{
					"error": "invalid data",
				})
				return
			}
			err, credential_uuid := credential.CreateCredential(
				reqBody["username"].(string),
				reqBody["password"].(string),
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
				"credential_data": "credential " + credential_uuid + " created successfully",
			})
			return
		case "GET":
			queryParamUsername := r.URL.Query().Get("username")
			queryParamPassword := r.URL.Query().Get("password")
			if (queryParamUsername == "") && (queryParamPassword == "") {
				Responser(w, r, true, 200, map[string]interface{}{
					"credential_data": credential.GetAllCredentials(),
				})
				return
			} else {
				Responser(w, r, true, 200, map[string]interface{}{
					"credential_data": credential.GetCredential(queryParamUsername, queryParamPassword),
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

