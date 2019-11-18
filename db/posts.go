package db

import "time"

// Post represents single post written by user.
type Post struct {
	ID        int
	Author    User
	Body      string
	CreatedOn time.Time
}

// NewPost returns new Post written by given user.
func NewPost(author User, body string) Post {
	return Post{
		Author:    author,
		Body:      body,
		CreatedOn: time.Now(),
	}
}
