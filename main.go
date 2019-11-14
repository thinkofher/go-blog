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

	dbwrapper, err := db.NewDBWrapper(CONFIG)
	if err != nil {
		panic(err)
	}
	defer dbwrapper.DB.Close()

	log.Println("Connection to database succesfull!")

	fs := http.StripPrefix("/static/", http.FileServer(http.Dir("./static/")))

	r := mux.NewRouter()

	r.Path("/").Handler(app.NewLoginHandler(dbwrapper, store)).Methods("GET", "POST")
	r.Path("/login").Handler(app.NewLoginHandler(dbwrapper, store)).Methods("GET", "POST")
	r.Path("/register").Handler(app.NewRegisterHandler(dbwrapper, store)).Methods("GET", "POST")
	r.Path("/logout").HandlerFunc(app.Logout(store)).Methods("GET")
	r.PathPrefix("/static/").Handler(fs).Methods("GET")

	authRouter := r.PathPrefix("/index").Subrouter()

	authRouter.Path("").Handler(app.NewIndexHandler(store)).Methods("GET")
	authRouter.Use(app.AuthenticationMiddleware(store))

	log.Fatal(http.ListenAndServe(":8080", r))
}
