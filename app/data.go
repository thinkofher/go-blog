package app

import (
	"github.com/thinkofher/go-blog/db"
)

type PageData struct {
	Title   string
	User    db.PublicUserData
	Flashes []interface{}
}

func NewPageData(title string) *PageData {
	return &PageData{
		Title: title,
	}
}

func (p *PageData) SetFlashes(flashes []interface{}) {
	p.Flashes = flashes
}

func (p *PageData) SetUserData(user db.PublicUserData) {
	p.User = user
}
