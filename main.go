package main

import (
	"encoding/gob"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
	_ "github.com/lib/pq"

	"github.com/thinkofher/go-blog/app"
	"github.com/thinkofher/go-blog/app/index"
	"github.com/thinkofher/go-blog/app/login"
	"github.com/thinkofher/go-blog/app/posts"
	"github.com/thinkofher/go-blog/app/register"
	"github.com/thinkofher/go-blog/db"
)

var store *sessions.CookieStore

func init() {
	store = sessions.NewCookieStore([]byte(os.Getenv("SESSION_KEY")))

	// Register PublicUserData to use in store
	gob.Register(db.PublicUserData{})
}

func main() {

	dbwrapper, err := db.NewWrapper(DBCONFIG)
	if err != nil {
		panic(err)
	}
	defer dbwrapper.DB.Close()

	log.Println("Connection to database succesfull!")

	fs := http.StripPrefix("/static/",
		http.FileServer(http.Dir("./static/")))

	r := mux.NewRouter()
	r.PathPrefix("/static/").Handler(fs).Methods("GET")
	r.Path("/logout").HandlerFunc(app.Logout(store, APPCONFIG)).Methods("GET")

	nonUsers := r.PathPrefix("").Subrouter()

	loginHandler := login.NewHandler(dbwrapper, store, APPCONFIG)
	nonUsers.Path("/").Handler(loginHandler).Methods("GET", "POST")
	nonUsers.Path("/login").Handler(loginHandler).Methods("GET", "POST")
	nonUsers.Path("/register").Handler(
		register.NewHandler(dbwrapper, store, APPCONFIG)).Methods("GET", "POST")
	nonUsers.Use(app.NonUsersOnly(store, APPCONFIG))

	auth := r.PathPrefix("/index").Subrouter()

	auth.Path("").Handler(index.NewHandler(dbwrapper, store, APPCONFIG)).Methods("GET")
	auth.Use(app.AuthenticationMiddleware(store, APPCONFIG))

	postsHandlers := r.PathPrefix("/post").Subrouter()
	postsHandlers.Path("/new").HandlerFunc(posts.NewPost(dbwrapper, store, APPCONFIG)).Methods("POST")
	postsHandlers.Use(app.AuthenticationMiddleware(store, APPCONFIG))

	log.Println("Starting application at port ':8080'.")
	log.Fatal(http.ListenAndServe(":8080", r))
}
