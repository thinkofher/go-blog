package app

import (
	"net/http"

	"github.com/gorilla/sessions"
)

func Logout(store *sessions.CookieStore) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}
