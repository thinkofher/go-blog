package main

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/thinkofher/go-blog/app"
	"github.com/thinkofher/go-blog/app/index"
	"github.com/thinkofher/go-blog/app/login"
	"github.com/thinkofher/go-blog/app/posts"
	"github.com/thinkofher/go-blog/app/register"
	"github.com/thinkofher/go-blog/app/user"
)

func router() *mux.Router {

	fs := http.StripPrefix("/static/",
		http.FileServer(http.Dir("./static/")))

	r := mux.NewRouter()
	r.
		PathPrefix("/static/").
		Handler(fs).
		Methods("GET")
	r.
		Path("/logout").
		HandlerFunc(app.Logout(store, APPCONFIG)).
		Methods("GET")

	nonUsers := r.PathPrefix("").Subrouter()
	nonUsers.
		Path("/").
		HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			http.Redirect(w, r, "/login", http.StatusFound)
		})
	nonUsers.
		Path("/login").
		Handler(login.NewHandler(dbwrapper, store, APPCONFIG)).
		Methods("GET", "POST")
	nonUsers.
		Path("/register").
		Handler(register.NewHandler(dbwrapper, store, APPCONFIG)).
		Methods("GET", "POST")
	nonUsers.
		Use(app.NonUsersOnly(store, APPCONFIG))

	indexHandlers := r.PathPrefix("/index").Subrouter()
	indexHandlers.
		Path("").
		Handler(index.NewHandler(dbwrapper, store, APPCONFIG)).
		Methods("GET")
	indexHandlers.
		Use(app.AuthenticationMiddleware(store, APPCONFIG))

	postsHandlers := r.PathPrefix("/post").Subrouter()
	postsHandlers.
		Path("/new").
		HandlerFunc(posts.NewPost(dbwrapper, store, APPCONFIG)).
		Methods("POST")
	postsHandlers.
		Path("/delete/{id:[0-9]+}").
		HandlerFunc(posts.RemovePost(dbwrapper, store, APPCONFIG)).
		Methods("POST", "GET")
	postsHandlers.
		Path("/edit/{id:[0-9]+}").
		HandlerFunc(posts.EditPost(dbwrapper, store, APPCONFIG)).
		Methods("POST")
	postsHandlers.
		Use(app.AuthenticationMiddleware(store, APPCONFIG))

	usersHandlers := r.PathPrefix("/user").Subrouter()
	usersHandlers.
		Path("/{id:[0-9]+}").
		Handler(user.NewHandler(dbwrapper, store, APPCONFIG)).
		Methods("GET")
	usersHandlers.
		Use(app.AuthenticationMiddleware(store, APPCONFIG))
	usersHandlers.Path("/upload/avatar").
		HandlerFunc(user.UploadAvatar(dbwrapper, store, APPCONFIG)).
		Methods("POST")

	return r
}
