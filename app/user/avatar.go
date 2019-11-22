package user

import (
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"

	"github.com/dchest/uniuri"
	"github.com/gorilla/sessions"
	"github.com/thinkofher/go-blog/app/utils"
	"github.com/thinkofher/go-blog/db"
)

var extensions = map[string]interface{}{
	".png":  struct{}{},
	".jpg":  struct{}{},
	".jpeg": struct{}{},
	".bmp":  struct{}{},
}

// UploadAvatar handles uploading avatars.
func UploadAvatar(client DBClient, store *sessions.CookieStore, config utils.AppConfig) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		session, err := store.Get(r, config.SessionName)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		userCookie, ok := session.Values[config.UserCookieKey].(db.PublicUserData)
		if !ok {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		r.ParseMultipartForm(10 << 20)

		file, header, err := r.FormFile("avatar")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer file.Close()

		_, ok = extensions[filepath.Ext(header.Filename)]
		if !ok {
			session.AddFlash("You can upload only image files!")
			err = session.Save(r, w)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			http.Redirect(w, r, "/index", http.StatusFound)
			return
		}

		randomFilename := uniuri.New() + filepath.Ext(header.Filename)
		avatar, err := os.Create(
			filepath.Join("static", "images", randomFilename))
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer avatar.Close()

		fileBytes, err := ioutil.ReadAll(file)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		_, err = avatar.Write(fileBytes)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		user, err := client.GetUserByID(userCookie.ID)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		err = removeOldAvatar(user.Avatar)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		err = client.UpdateAvatar(userCookie.ID, randomFilename)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		session.AddFlash("Your avatar has been updated!")
		err = session.Save(r, w)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		http.Redirect(w, r, "/index", http.StatusFound)
	}
}

func removeOldAvatar(filename string) error {
	if filename == db.DefaultAvatar {
		return nil
	}

	return os.Remove(filepath.Join("static", "images", filename))
}
