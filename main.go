package main

import (
	"encoding/gob"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/sessions"
	_ "github.com/lib/pq"

	"github.com/thinkofher/go-blog/db"
)

var store *sessions.CookieStore
var dbwrapper db.Wrapper
var err error

func init() {
	store = sessions.NewCookieStore([]byte(os.Getenv("SESSION_KEY")))

	// Register PublicUserData to use in store
	gob.Register(db.PublicUserData{})
}

func main() {

	dbwrapper, err = db.NewWrapper(DBCONFIG)
	if err != nil {
		panic(err)
	}
	defer dbwrapper.DB.Close()

	log.Println("Connection to database succesfull!")

	r := router()

	log.Println("Starting application at port ':8080'.")
	log.Fatal(http.ListenAndServe(":8080", r))
}
