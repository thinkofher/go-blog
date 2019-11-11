package register

import "github.com/thinkofher/go-blog/db"

// DBClient handles connection between app and database
// for registering purposes.
type DBClient interface {
	// SetUser registers new user in database or returns
	// error if operation failed.
	SetUser(db.User) error
}
