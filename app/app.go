package app

import (
	"log"
	"net/http"

	"github.com/thinkofher/go-blog/app/login"
	"github.com/thinkofher/go-blog/app/register"
	"github.com/thinkofher/go-blog/db"
)

type loginHandler struct {
	tmpl BlogTemplate
	db   login.DBClient
}

func NewLoginHandler(db login.DBClient) *loginHandler {
	return &loginHandler{
		tmpl: NewBlogTemplate("login", *NewPageData("Login")),
		db:   db,
	}
}

func (h loginHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	tmpl, err := h.tmpl.Template()
	if err != nil {
		log.Fatal(err)
	}

	if err = tmpl.ExecuteTemplate(w, templatesBase, h.tmpl.TemplateData()); err != nil {
		log.Fatal("Could not execute login templates.")
	}
}

type registerHandler struct {
	tmpl BlogTemplate
	db   register.DBClient
}

func NewRegisterHandler(db register.DBClient) *registerHandler {
	return &registerHandler{
		tmpl: NewBlogTemplate("register", *NewPageData("Register")),
		db:   db,
	}
}

func (h registerHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	tmpl, err := h.tmpl.Template()
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Register method: ", r.Method)
	if r.Method == http.MethodPost {
		// err := r.ParseForm()
		if err != nil {
			log.Println(err.Error())
		}
		// TODO: handle hashing password error
		log.Println("Creating user...")
		user, _ := db.NewUser(
			r.FormValue("username"),
			r.FormValue("password"),
			r.FormValue("email"))

		err = h.db.SetUser(user)
		if err != nil {
			log.Println("Could not register user with username: ", user.Username)
			http.Redirect(w, r, "/login", http.StatusSeeOther)
		}
		log.Println("User: ", user.Username, " created.")
		http.Redirect(w, r, "/login", http.StatusSeeOther)
	} else {
		if err = tmpl.ExecuteTemplate(w, templatesBase, h.tmpl.TemplateData()); err != nil {
			log.Fatal("Could not execute register templates.")
		}
	}
}
