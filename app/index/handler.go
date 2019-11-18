package index

import (
	"log"
	"net/http"

	"github.com/gorilla/sessions"

	"github.com/thinkofher/go-blog/app/blog"
	"github.com/thinkofher/go-blog/app/posts"
	"github.com/thinkofher/go-blog/app/utils"
	"github.com/thinkofher/go-blog/db"
)

type handler struct {
	tmpl   blog.Renderer
	store  *sessions.CookieStore
	config utils.AppConfig
	db     posts.DBClient
}

// NewHandler returns Handler for index page.
func NewHandler(db posts.DBClient, store *sessions.CookieStore, config utils.AppConfig) http.Handler {
	return &handler{
		tmpl:   blog.NewRenderer("index", *blog.NewData("Index")),
		db:     db,
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

	posts, err := h.db.GetPosts()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	data.Posts = posts

	userCookie, ok := session.Values[h.config.UserCookieKey].(db.PublicUserData)
	if !ok {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	data.SetUserData(userCookie)

	err = session.Save(r, w)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err = tmpl.ExecuteTemplate(w, blog.TemplatesBase, data); err != nil {
		log.Fatal("Could not execute index templates.")
	}
}
