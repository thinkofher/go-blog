package app

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
)

// AuthenticationMiddleware returns MiddlewareFunc, which uses
// given CookieStore to authenticate is user is logged in.
func AuthenticationMiddleware(store *sessions.CookieStore) mux.MiddlewareFunc {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			session, err := store.Get(r, SessionName)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			_, ok := session.Values[userCookieKey]

			if !ok {
				session.AddFlash("You have to be logged in.")
				err = session.Save(r, w)
				if err != nil {
					http.Error(w, err.Error(), http.StatusInternalServerError)
					return
				}
				http.Redirect(w, r, "/login", http.StatusFound)
				return
			}
			next.ServeHTTP(w, r)
		})
	}
}

// NonUsersOnly returns mux.MiddlewareFunc, which uses
// given CookieStore to authenticate is user is logged out.
func NonUsersOnly(store *sessions.CookieStore) mux.MiddlewareFunc {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			session, err := store.Get(r, SessionName)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			_, ok := session.Values[userCookieKey]

			if ok {
				session.AddFlash("You have to be logged out.")
				err = session.Save(r, w)
				if err != nil {
					http.Error(w, err.Error(), http.StatusInternalServerError)
					return
				}
				http.Redirect(w, r, "/index", http.StatusFound)
				return
			}
			next.ServeHTTP(w, r)
		})
	}
}
