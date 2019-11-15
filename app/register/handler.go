package register

import (
	"log"
	"net/http"

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
		tmpl:   blog.NewRenderer("register", *blog.NewData("Register")),
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

	if r.Method == http.MethodPost {
		// err := r.ParseForm()
		if err != nil {
			log.Println(err.Error())
		}

		user, _ := db.NewUser(
			r.FormValue("username"),
			r.FormValue("password"),
			r.FormValue("email"))

		err = h.db.SetUser(user)
		if err != nil {
			session.AddFlash("Username or Email are already taken.")

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
		log.Fatal("Could not execute register templates.")
	}
}
