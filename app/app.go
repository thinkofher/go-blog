package app

import (
	"fmt"
	"html/template"
	"log"
	"net/http"

	"github.com/thinkofher/go-blog/app/login"
	"github.com/thinkofher/go-blog/app/register"
)

type PageInfo struct {
	Title string
}

func NewPageInfo(title string) *PageInfo {
	return &PageInfo{
		Title: title,
	}
}

const baseTemplate string = "base"

func loadBaseTemplate() (*template.Template, error) {
	return template.ParseGlob("templates/*.tmpl")
}

func loadTemplates(t *template.Template, foldername string) (*template.Template, error) {
	return t.ParseGlob(fmt.Sprintf("templates/%s/*.tmpl", foldername))
}

func newTemplate(foldername string) (*template.Template, error) {
	tmpl, err := loadBaseTemplate()
	if err != nil {
		return nil, err
	}

	tmpl, err = loadTemplates(tmpl, foldername)
	if err != nil {
		return nil, err
	}

	return tmpl, nil
}

type loginHandler struct {
	Info PageInfo
	db   *login.DBClient
}

func NewLoginHandler() *loginHandler {
	return &loginHandler{
		Info: *NewPageInfo("Login"),
	}
}

func (h loginHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	tmpl, err := newTemplate("login")
	if err != nil {
		log.Fatal(err)
	}

	if err = tmpl.ExecuteTemplate(w, baseTemplate, h.Info); err != nil {
		log.Fatal("Could not execute login templates.")
	}
}

type registerHandler struct {
	Info PageInfo
	db   *register.DBClient
}

func NewRegisterHandler() *registerHandler {
	return &registerHandler{
		Info: *NewPageInfo("Register"),
	}
}

func (h registerHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	tmpl, err := newTemplate("register")
	if err != nil {
		log.Fatal(err)
	}

	if err = tmpl.ExecuteTemplate(w, baseTemplate, h.Info); err != nil {
		log.Fatal("Could not execute register templates.")
	}
}
