package user

import "github.com/thinkofher/go-blog/db"

// DBClient handles connection between app and database
// for purposes of user package.
type DBClient interface {
	// GetPostsByUser method returns all posts written
	// by user with given id.
	GetPostsByUser(userid int) ([]db.Post, error)
	// Returns complete database representation of
	// User with given id.
	GetUserByID(id int) (db.User, error)
}
