package app

import (
	"log"
	"net/http"

	"golang.org/x/crypto/bcrypt"

	"github.com/gorilla/sessions"
	"github.com/thinkofher/go-blog/app/login"
)

var userCookieKey = "user-cookie"

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

	if r.Method == http.MethodPost {
		username := r.FormValue("username")
		password := r.FormValue("password")

		fullUserData, err := h.db.GetUser(username)

		if err != nil {
			session.AddFlash("There is no user with given username.")
			err = session.Save(r, w)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			http.Redirect(w, r, "/login", http.StatusFound)
			return
		}

		if err = bcrypt.CompareHashAndPassword(
			fullUserData.HashedPassword, []byte(password)); err != nil {

			session.AddFlash("Incorrect password.")
			err = session.Save(r, w)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			http.Redirect(w, r, "/login", http.StatusFound)
			return
		}

		userCookie := fullUserData.ToPublicUserData()

		session.Values[userCookieKey] = userCookie
		err = session.Save(r, w)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		http.Redirect(w, r, "/index", http.StatusFound)
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
