package app

import (
	"log"
	"net/http"

	"github.com/thinkofher/go-blog/app/login"
	"github.com/thinkofher/go-blog/app/register"
)

type loginHandler struct {
	tmpl BlogTemplate
	db   login.DBClient
}

func NewLoginHandler() *loginHandler {
	return &loginHandler{
		tmpl: NewBlogTemplate("login", *NewPageData("Login")),
	}
}

func (h loginHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	tmpl, err := h.tmpl.Template()
	if err != nil {
		log.Fatal(err)
	}

	if err = tmpl.ExecuteTemplate(w, templatesBase, h.tmpl.TemplateData()); err != nil {
		log.Fatal("Could not execute login templates.")
	}
}

type registerHandler struct {
	tmpl BlogTemplate
	db   register.DBClient
}

func NewRegisterHandler() *registerHandler {
	return &registerHandler{
		tmpl: NewBlogTemplate("register", *NewPageData("Register")),
	}
}

func (h registerHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	tmpl, err := h.tmpl.Template()
	if err != nil {
		log.Fatal(err)
	}

	if err = tmpl.ExecuteTemplate(w, templatesBase, h.tmpl.TemplateData()); err != nil {
		log.Fatal("Could not execute register templates.")
	}
}
