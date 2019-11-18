package blog

import (
	"github.com/thinkofher/go-blog/db"
)

// Data represents data for client
// to render on the page.
type Data struct {
	Title   string
	User    db.PublicUserData
	Flashes []interface{}
	Posts   []db.Post
}

// NewData returns pointer to PageData with
// with given title as Title of the page.
func NewData(title string) *Data {
	return &Data{
		Title: title,
	}
}

// SetFlashes methods accepts list of flashes to view
// them next time to the user.
func (p *Data) SetFlashes(flashes []interface{}) {
	p.Flashes = flashes
}

// SetUserData methods accepts PublicUserData to render
// user specific content.
func (p *Data) SetUserData(user db.PublicUserData) {
	p.User = user
}
