package login

import (
	"net/http"

	"github.com/thinkofher/go-blog/db"
)

// DBClient handles connection between app and database
// for logging purposes.
type DBClient interface {
	// GetUser returns User data from database
	// under given id number.
	GetUser(username string) db.User
}

// Middleware used for user authentication,
// implements gorilla/mux MiddlewareFunc.
type Middleware interface {
	Authenticate(http.Handler) http.Handler
}
