package index

import (
	"log"
	"net/http"

	"github.com/gorilla/sessions"

	"github.com/thinkofher/go-blog/app/blog"
	"github.com/thinkofher/go-blog/app/utils"
)

type handler struct {
	tmpl   blog.Renderer
	store  *sessions.CookieStore
	config utils.AppConfig
}

// NewHandler returns Handler for index page.
func NewHandler(store *sessions.CookieStore, config utils.AppConfig) http.Handler {
	return &handler{
		tmpl:   blog.NewRenderer("index", *blog.NewData("Index")),
		store:  store,
		config: config,
	}
}

// ServeHTTP satisfies http.Handler interface.
func (h handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	tmpl, err := h.tmpl.Render()
	if err != nil {
		log.Fatal(err)
	}

	session, err := h.store.Get(r, h.config.SessionName)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	data := h.tmpl.Data()
	if flashes := session.Flashes(); len(flashes) > 0 {
		data.SetFlashes(flashes)
	}

	err = session.Save(r, w)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err = tmpl.ExecuteTemplate(w, blog.TemplatesBase, data); err != nil {
		log.Fatal("Could not execute login templates.")
	}
}
