package app

import (
	"github.com/thinkofher/go-blog/db"
)

// PageData represents data for client
// to render on the page.
type PageData struct {
	Title   string
	User    db.PublicUserData
	Flashes []interface{}
}

// NewPageData returns pointer to PageData with
// with given title as Title of the page.
func NewPageData(title string) *PageData {
	return &PageData{
		Title: title,
	}
}

// SetFlashes methods accepts list of flashes to view
// them next time to the user.
func (p *PageData) SetFlashes(flashes []interface{}) {
	p.Flashes = flashes
}

// SetUserData methods accepts PublicUserData to render
// user specific content.
func (p *PageData) SetUserData(user db.PublicUserData) {
	p.User = user
}
