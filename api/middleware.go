package api

import (
	"net/http"

	"github.com/gorilla/sessions"

	"secretary/alpha/internal"
)

var store = sessions.NewCookieStore([]byte("my_secret_key"))

func setHeaders(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
}

func isAuthenticated(r *http.Request) (interface{}, *internal.User) {
	// FIXME change the secret
	user := internal.User{}
	session, _ := store.Get(r, "session.id")
	return session.Values["authenticated"], user.GetUser(session.Values["username"].(string))
}

func Middleware(w http.ResponseWriter, r *http.Request, secure ...bool) bool {
	if (len(secure) > 0) && (!secure[0]) {
		setHeaders(w)
		return true
	}

	authenticated, _ := isAuthenticated(r)
	if (authenticated != nil) && (authenticated != false) {
		setHeaders(w)
		return true
	} else {
		return false
	}
}
