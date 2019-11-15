package app

import (
	"log"
	"net/http"

	"github.com/gorilla/sessions"
)

type indexHandler struct {
	tmpl  BlogTemplate
	store *sessions.CookieStore
}

// NewIndexHandler returns Handler for index page.
func NewIndexHandler(store *sessions.CookieStore) http.Handler {
	return &indexHandler{
		tmpl:  NewBlogTemplate("index", *NewPageData("Index")),
		store: store,
	}
}

func (h indexHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
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
