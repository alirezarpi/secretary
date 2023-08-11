package api

import (
	"net/http"

	"github.com/gorilla/sessions"

	"secretary/alpha/internal"
	"secretary/alpha/utils"
)

// FIXME change the secret
var store = sessions.NewCookieStore([]byte("my_secret_key"))

func setHeaders(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
}

func isAuthenticated(r *http.Request) (interface{}, *internal.User) {
	user := internal.User{}
	session, err := store.Get(r, "session.id")
	if err != nil {
		utils.Logger("err", err.Error())
		return false, nil
	}
	if (len(session.Values) == 0) || (session.Values["username"] == nil) {
		return false, nil
	}
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
