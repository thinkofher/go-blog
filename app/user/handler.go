package user

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"

	"github.com/thinkofher/go-blog/app/blog"
	"github.com/thinkofher/go-blog/app/utils"
	"github.com/thinkofher/go-blog/db"
)

type handler struct {
	tmpl   blog.Renderer
	db     DBClient
	store  *sessions.CookieStore
	config utils.AppConfig
}

// NewHandler returns Handler for register page.
func NewHandler(db DBClient, store *sessions.CookieStore, config utils.AppConfig) http.Handler {
	return &handler{
		// TODO: Create own templates
		tmpl:   blog.NewRenderer("user", *blog.NewData("User")),
		db:     db,
		store:  store,
		config: config,
	}
}

// ServeHTTP satisfies http.Handler interface.
func (h handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	tmpl, err := h.tmpl.Render()
	if err != nil {
		log.Fatal(err)
	}

	session, err := h.store.Get(r, h.config.SessionName)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	id, ok := vars["id"]
	if !ok {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	data := h.tmpl.Data()
	if flashes := session.Flashes(); len(flashes) > 0 {
		data.SetFlashes(flashes)
	}

	intid, _ := strconv.Atoi(id)

	posts, err := h.db.GetPostsByUser(intid)
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

	requestedUser, err := h.db.GetUserByID(intid)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	data.SetTargetUser(requestedUser)
	data.Title = requestedUser.Username

	err = session.Save(r, w)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err = tmpl.ExecuteTemplate(w, blog.TemplatesBase, data); err != nil {
		log.Fatal("Could not execute register templates.")
	}
}
