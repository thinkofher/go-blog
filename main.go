package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
	"github.com/thinkofher/go-blog/app"
)

func main() {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err)
	}
	log.Println("Connection to database succesfull!")

	fs := http.StripPrefix("/static/", http.FileServer(http.Dir("./static/")))

	r := mux.NewRouter()

	r.Path("/").Handler(app.NewLoginHandler()).Methods("GET", "POST")
	r.Path("/login").Handler(app.NewLoginHandler()).Methods("GET", "POST")
	r.Path("/register").Handler(app.NewRegisterHandler()).Methods("GET", "POST")
	r.PathPrefix("/static/").Handler(fs).Methods("GET")

	log.Fatal(http.ListenAndServe(":8080", r))
}
