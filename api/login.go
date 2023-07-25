package api

import (
	"log"
	"net/http"

	"github.com/gorilla/sessions"

	"secretary/alpha/utils"
	"secretary/alpha/internal"
)

func LoginAPI(w http.ResponseWriter, r *http.Request) {
	Middleware(w, r, false)

	if r.Method != "POST" {
		Responser(w, r, false, 405, map[string]interface{}{
			"message": "method not allowed",
		})
	}

	// FIXME use ENV for secretkey
	var store = sessions.NewCookieStore([]byte("my_secret_key"))
	reqBody, err := utils.HandleReqJson(r)
	if err != nil {
		log.Println(err)
	}

	user := &internal.User{}
	retrievedUser := user.GetUser(reqBody["username"].(string))
	if retrievedUser == nil {
		Responser(w, r, false, 401, map[string]interface{}{
			"message": "Unauthorized",
		})
	}
	if retrievedUser.CheckPassword(reqBody["password"].(string)) {
		session, _ := store.Get(r, "session.id")
		session.Values["authenticated"] = true
		store.MaxAge(86400 * 3)
		session.Save(r, w)
		Responser(w, r, true, 200, map[string]interface{}{
			"message": "login successfully",
		})
	} else {
		Responser(w, r, false, 401, map[string]interface{}{
			"message": "Unauthorized",
		})
	}
}
