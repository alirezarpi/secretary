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

	username := reqBody["username"].(string)
	password := reqBody["password"].(string)

	user := &internal.User{}
	result := user.GetUser(username)
	if result == nil {
		Responser(w, r, false, 401, map[string]interface{}{
			"message": "Unauthorized",
		})
	}
	session, _ := store.Get(r, "session.id")
	if user.CheckPassword(password) {
		session.Values["authenticated"] = true
		session.Save(r, w)
	} else {
		Responser(w, r, false, 401, map[string]interface{}{
			"message": "Unauthorized",
		})
	}
	Responser(w, r, true, 200, map[string]interface{}{
		"message": "login successfully",
	})
}
