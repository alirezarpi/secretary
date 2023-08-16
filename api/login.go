package api

import (
	"net/http"
	"time"

	"secretary/alpha/internal"
	"secretary/alpha/internal/constants"
	"secretary/alpha/utils"
)

func LoginAPI(w http.ResponseWriter, r *http.Request) {
	Middleware(w, r, false)

	if r.Method != "POST" {
		Responser(w, r, false, 405, map[string]interface{}{
			"message": "method not allowed",
		})
		return
	}

	reqBody, err := utils.HandleReqJson(r)
	if err != nil {
		utils.Logger("err", err.Error())
	}
	retrievedUser := &internal.User{}
	retrievedUser = retrievedUser.GetUser(reqBody["username"].(string))
	if retrievedUser == nil {
		Responser(w, r, false, 401, map[string]interface{}{
			"message": "unauthorized",
		})
		return
	}
	if retrievedUser.CheckPassword(reqBody["password"].(string)) {
		session, _ := store.Get(r, "sc_session_id")
		session.Values["sc_authenticated"] = true
		session.Values["sc_username"] = reqBody["username"]
		session.Values["sc_time"] = time.Now().Unix()
		session.Options.MaxAge = constants.HTTP_SC_MAXAGE
		store.Save(r, w, session)
		Responser(w, r, true, 200, map[string]interface{}{
			"message": "login successfully",
		})
		return
	} else {
		Responser(w, r, false, 401, map[string]interface{}{
			"message": "unauthorized",
		})
		return
	}
}
