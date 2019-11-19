package posts

import (
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
	"github.com/thinkofher/go-blog/app/utils"
	"github.com/thinkofher/go-blog/db"
)

// DBClient handles connection between app and database
// to manipulate users posts.
type DBClient interface {
	// GetUser returns User data from database
	// under given id number.

	SetPost(post db.Post) error
	GetPosts() ([]db.Post, error)
	RemovePost(postID int, authorID int) error
	// GetUserPosts(username string) ([]db.Post, error)
}

// NewPost handles creating blog posts.
func NewPost(client DBClient, store *sessions.CookieStore, config utils.AppConfig) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		postBody := r.FormValue("post-body")

		session, err := store.Get(r, config.SessionName)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		blankUser := db.User{}
		userCookie, ok := session.Values[config.UserCookieKey].(db.PublicUserData)
		if !ok {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		blankUser.ID = userCookie.ID
		post := db.NewPost(blankUser, postBody)

		err = client.SetPost(post)
		if err != nil {
			session.AddFlash("Could not send post.")

			err = session.Save(r, w)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			http.Redirect(w, r, "/index", http.StatusFound)
			return
		}

		session.AddFlash("Your post has been sended.")
		err = session.Save(r, w)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		http.Redirect(w, r, "/index", http.StatusFound)
	}
}

// RemovePost handles removing posts.
func RemovePost(client DBClient, store *sessions.CookieStore, config utils.AppConfig) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)

		session, err := store.Get(r, config.SessionName)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		id, ok := vars["id"]
		if !ok {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		userCookie, ok := session.Values[config.UserCookieKey].(db.PublicUserData)
		if !ok {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		idint, _ := strconv.Atoi(id)
		err = client.RemovePost(idint, userCookie.ID)
		if err != nil {
			// TODO: Do not show this to user!
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		session.AddFlash("Post has been deleted.")
		err = session.Save(r, w)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		http.Redirect(w, r, "/index", http.StatusFound)
	}
}
