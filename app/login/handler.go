package login

import (
	"log"
	"net/http"

	"github.com/gorilla/sessions"
	"golang.org/x/crypto/bcrypt"

	"github.com/thinkofher/go-blog/app/blog"
	"github.com/thinkofher/go-blog/app/utils"
)

type handler struct {
	tmpl   blog.Renderer
	db     DBClient
	store  *sessions.CookieStore
	config utils.AppConfig
}

// NewHandler returns Handler for login page.
func NewHandler(db DBClient, store *sessions.CookieStore, config utils.AppConfig) http.Handler {
	return &handler{
		tmpl:   blog.NewRenderer("login", *blog.NewData("Login")),
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

		session.Values[h.config.UserCookieKey] = userCookie
		err = session.Save(r, w)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		http.Redirect(w, r, "/index", http.StatusFound)
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
