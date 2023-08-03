package api

import (
	"log"
	"net/http"

	"github.com/gorilla/sessions"

	"secretary/alpha/internal"
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

	// FIXME use ENV for secretkey
	var store = sessions.NewCookieStore([]byte("my_secret_key"))
	reqBody, err := utils.HandleReqJson(r)
	if err != nil {
		log.Println(err)
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
		session, _ := store.Get(r, "session.id")
		session.Values["authenticated"] = true
		session.Values["username"] = reqBody["username"]
		store.MaxAge(2592000)
		session.Save(r, w)
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
