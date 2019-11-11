package db

// User represents single user data from database.
type User struct {
	// ID is unique to every user.
	ID             int
	Username       string
	HashedPassword []byte
}
