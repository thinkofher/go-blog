package app

import (
	"log"
	"net/http"

	"github.com/gorilla/sessions"
	"github.com/thinkofher/go-blog/app/login"
	"github.com/thinkofher/go-blog/app/register"
	"github.com/thinkofher/go-blog/db"
)

var SessionName = "goblog-session"

type loginHandler struct {
	tmpl  BlogTemplate
	db    login.DBClient
	store *sessions.CookieStore
}

func NewLoginHandler(db login.DBClient, store *sessions.CookieStore) *loginHandler {
	return &loginHandler{
		tmpl:  NewBlogTemplate("login", *NewPageData("Login")),
		db:    db,
		store: store,
	}
}

func (h loginHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	tmpl, err := h.tmpl.Template()
	if err != nil {
		log.Fatal(err)
	}

	session, err := h.store.Get(r, SessionName)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	data := h.tmpl.TemplateData()
	if flashes := session.Flashes(); len(flashes) > 0 {
		data.SetFlashes(flashes)
	}

	err = session.Save(r, w)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err = tmpl.ExecuteTemplate(w, templatesBase, data); err != nil {
		log.Fatal("Could not execute login templates.")
	}
}

type registerHandler struct {
	tmpl  BlogTemplate
	db    register.DBClient
	store *sessions.CookieStore
}

func NewRegisterHandler(db register.DBClient, store *sessions.CookieStore) *registerHandler {
	return &registerHandler{
		tmpl:  NewBlogTemplate("register", *NewPageData("Register")),
		db:    db,
		store: store,
	}
}

func (h registerHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	tmpl, err := h.tmpl.Template()
	if err != nil {
		log.Fatal(err)
	}

	session, err := h.store.Get(r, SessionName)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if r.Method == http.MethodPost {
		// err := r.ParseForm()
		if err != nil {
			log.Println(err.Error())
		}

		// TODO: handle hashing password error
		user, _ := db.NewUser(
			r.FormValue("username"),
			r.FormValue("password"),
			r.FormValue("email"))

		err = h.db.SetUser(user)
		if err != nil {
			session.AddFlash("Username or Password are already taken.")

			err = session.Save(r, w)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			http.Redirect(w, r, "/register", http.StatusFound)
			return
		}

		session.AddFlash("Your account has been created.")
		err = session.Save(r, w)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		http.Redirect(w, r, "/login", http.StatusFound)
		return
	}

	data := h.tmpl.TemplateData()
	if flashes := session.Flashes(); len(flashes) > 0 {
		data.SetFlashes(flashes)
	}

	err = session.Save(r, w)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err = tmpl.ExecuteTemplate(w, templatesBase, data); err != nil {
		log.Fatal("Could not execute register templates.")
	}
}
