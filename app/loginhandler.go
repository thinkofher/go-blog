package app

import (
	"log"
	"net/http"

	"github.com/gorilla/sessions"
	"github.com/thinkofher/go-blog/app/login"
)

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
