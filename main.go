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
	"github.com/thinkofher/go-blog/db"
)

var store *sessions.CookieStore

func init() {
	store = sessions.NewCookieStore([]byte(os.Getenv("SESSION_KEY")))

	// Register PublicUserData to use in store
	gob.Register(db.PublicUserData{})
}

func main() {

	dbwrapper, err := db.NewWrapper(CONFIG)
	if err != nil {
		panic(err)
	}
	defer dbwrapper.DB.Close()

	log.Println("Connection to database succesfull!")

	fs := http.StripPrefix("/static/",
		http.FileServer(http.Dir("./static/")))

	r := mux.NewRouter()
	r.PathPrefix("/static/").Handler(fs).Methods("GET")
	r.Path("/logout").HandlerFunc(app.Logout(store)).Methods("GET")

	nonUsers := r.PathPrefix("").Subrouter()

	login := app.NewLoginHandler(dbwrapper, store)
	nonUsers.Path("/").Handler(login).Methods("GET", "POST")
	nonUsers.Path("/login").Handler(login).Methods("GET", "POST")
	nonUsers.Path("/register").Handler(
		app.NewRegisterHandler(dbwrapper, store)).Methods("GET", "POST")
	nonUsers.Use(app.NonUsersOnly(store))

	auth := r.PathPrefix("/index").Subrouter()

	auth.Path("").Handler(app.NewIndexHandler(store)).Methods("GET")
	auth.Use(app.AuthenticationMiddleware(store))

	log.Println("Starting application at port ':8080'.")

	log.Fatal(http.ListenAndServe(":8080", r))
}
