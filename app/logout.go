package app

import (
	"net/http"

	"github.com/gorilla/sessions"
	"github.com/thinkofher/go-blog/app/utils"
)

// Logout clears cookies containing user data and
// redirect to the login page.
func Logout(store *sessions.CookieStore, config utils.AppConfig) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		session, err := store.Get(r, config.SessionName)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		_, ok := session.Values[config.UserCookieKey]
		if ok {
			delete(session.Values, config.UserCookieKey)
			err = session.Save(r, w)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
		}
		http.Redirect(w, r, "/login", http.StatusFound)
	}
}
