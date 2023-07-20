package api

import (
	"log"
	"net/http"

	"github.com/gorilla/sessions"

	"secretary/alpha/utils"
)

func LoginAPI(w http.ResponseWriter, r *http.Request) {
	Middleware(w, r, false)

    if (r.Method != "POST") {
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

    storedPassword, exists := users[username]
    if exists {
        // It returns a new session if the sessions doesn't exist
        session, _ := store.Get(r, "session.id")
        if storedPassword == password {
            session.Values["authenticated"] = true
            // Saves all sessions used during the current request
            session.Save(r, w)
        } else {
            http.Error(w, "Invalid Credentials", http.StatusUnauthorized)
        }
        w.Write([]byte("Login successfully!"))
    }

}
