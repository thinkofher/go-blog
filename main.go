package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
	"github.com/thinkofher/go-blog/app"
	"github.com/thinkofher/go-blog/db"
)

func main() {

	dbwrapper, err := db.NewDBWrapper(CONFIG)
	if err != nil {
		panic(err)
	}
	defer dbwrapper.DB.Close()

	log.Println("Connection to database succesfull!")

	fs := http.StripPrefix("/static/", http.FileServer(http.Dir("./static/")))

	r := mux.NewRouter()

	r.Path("/").Handler(app.NewLoginHandler()).Methods("GET", "POST")
	r.Path("/login").Handler(app.NewLoginHandler()).Methods("GET", "POST")
	r.Path("/register").Handler(app.NewRegisterHandler()).Methods("GET", "POST")
	r.PathPrefix("/static/").Handler(fs).Methods("GET")

	log.Fatal(http.ListenAndServe(":8080", r))
}
