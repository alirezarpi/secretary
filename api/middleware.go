package api

import (
	"net/http"
	"time"

	"github.com/gorilla/sessions"

	"secretary/alpha/internal"
	"secretary/alpha/internal/constants"
	"secretary/alpha/utils"
)

// FIXME change the secret
var store = sessions.NewCookieStore([]byte(constants.HTTP_SC_SECRET))

func setHeaders(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
}

func isAuthenticated(r *http.Request) (interface{}, *internal.User) {
	user := internal.User{}
	session, err := store.Get(r, "sc_session_id")
	if err != nil {
		return false, nil
	}
	if (len(session.Values) == 0) || (session.Values["sc_username"] == nil) {
		return false, nil
	}

	sessionCookieCreatedTime, ok := session.Values["sc_time"].(int64)
	if !ok {
		utils.Logger("err", "could not extract value sc_time from session-cookie")
		return false, nil
	}

	currentTime := time.Now().Unix()
	sessionCookieMaxSessionAge := int64(constants.HTTP_SC_MAXAGE)

	if (currentTime - sessionCookieCreatedTime) > sessionCookieMaxSessionAge {
		utils.Logger("debug", "session for user < "+session.Values["sc_username"].(string)+" > has been expired")
		return false, nil
	}
	return session.Values["sc_authenticated"], user.GetUser(session.Values["sc_username"].(string))
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
		Responser(w, r, false, 401, map[string]interface{}{
			"message": "unauthorized",
		})
		return false
	}
}
