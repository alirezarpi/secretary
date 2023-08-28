package api

import (
	"net/http"

	"secretary/alpha/internal/audit"
	"secretary/alpha/utils"
)

func LogoutAPI(w http.ResponseWriter, r *http.Request) {
	if Middleware(w, r) {
		session, err := store.Get(r, "sc_session_id")
		if err != nil {
			utils.Logger("err", err.Error())
			Responser(w, r, false, 500, map[string]interface{}{
				"error": "internal error",
			})
			return
		}
		session.Options.MaxAge = -1
		session.Save(r, w)
		Responser(w, r, true, 200, map[string]interface{}{
			"message": "you're logged-out",
		})
		audit.Audit("[user] [action:logout] user " + session.Values["sc_username"].(string) + " logged out.")
		return
	} else {
		Responser(w, r, false, 401, map[string]interface{}{
			"error": "unauthorized",
		})
		return
	}
}
