package api

import (
	"net/http"

	"github.com/gorilla/sessions"
)


func setHeaders(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
}

func isAuthenticated(r *http.Request) interface{} {
	// FIXME change the secret
	var store = sessions.NewCookieStore([]byte("my_secret_key"))
	session, _ := store.Get(r, "session.id")
    authenticated := session.Values["authenticated"]
	return authenticated
}

func Middleware(w http.ResponseWriter, r *http.Request, secure ...bool) bool {
	if (len(secure) > 0) && (!secure[0]) {
		setHeaders(w)
		return true
	} 

	authenticated := isAuthenticated(r)
	if (authenticated != nil) && (authenticated != false) {
		setHeaders(w)
		return true
    } else {
		return false
    }
}

