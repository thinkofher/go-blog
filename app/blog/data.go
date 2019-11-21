package blog

import (
	"github.com/thinkofher/go-blog/db"
)

// Data represents data for client
// to render on the page.
type Data struct {
	Title        string
	User         db.PublicUserData
	Flashes      []interface{}
	Posts        []db.Post
	PageSpecific PageSpecific
}

// PageSpecific represents data specific
// to page being viewed by user.
type PageSpecific struct {
	User db.PublicUserData
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

// SetTargetUser takes given user and use its
// public data to render when needed.
func (p *Data) SetTargetUser(user db.User) {
	p.PageSpecific.User = *user.ToPublicUserData()
}
