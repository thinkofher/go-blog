package app

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"

	"github.com/thinkofher/go-blog/app/utils"
)

// AuthenticationMiddleware returns MiddlewareFunc, which uses
// given CookieStore to authenticate user, if is he logged in.
func AuthenticationMiddleware(store *sessions.CookieStore, config utils.AppConfig) mux.MiddlewareFunc {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			session, err := store.Get(r, config.SessionName)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			_, ok := session.Values[config.UserCookieKey]

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
// given CookieStore to authenticate user, if is he logged out.
func NonUsersOnly(store *sessions.CookieStore, config utils.AppConfig) mux.MiddlewareFunc {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			session, err := store.Get(r, config.SessionName)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			_, ok := session.Values[config.UserCookieKey]

			if ok {
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
