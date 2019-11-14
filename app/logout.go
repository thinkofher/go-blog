package app

import (
	"net/http"

	"github.com/gorilla/sessions"
)

// Logout clears cookies containing user data and
// redirect to the login page.
func Logout(store *sessions.CookieStore) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		session, err := store.Get(r, SessionName)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		_, ok := session.Values[userCookieKey]
		if ok {
			delete(session.Values, userCookieKey)
			err = session.Save(r, w)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
		}
		http.Redirect(w, r, "/login", http.StatusFound)
	}
}
