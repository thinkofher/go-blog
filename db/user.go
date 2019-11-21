package db

import (
	"time"

	"golang.org/x/crypto/bcrypt"
)

// CRYPTOCOST contains value of cost specific for
// go-blog application.
const CRYPTOCOST = 14

// User represents single user data from database.
type User struct {
	// ID is unique to every user.
	ID             int
	Username       string
	HashedPassword []byte
	Email          string
	CreatedOn      time.Time
	LastLogin      time.Time
}

// NewUser constructor hashes given password string
// and returns User model struct. Returns empty User struct
// and error if hashing failes.
func NewUser(username, password, email string) (User, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword(
		[]byte(password), CRYPTOCOST)
	if err != nil {
		return User{}, err
	}

	return User{
		Username:       username,
		HashedPassword: hashedPassword,
		Email:          email,
		CreatedOn:      time.Now(),
		LastLogin:      time.Now(),
	}, nil
}

// ToPublicUserData transforms full User data
// from database to PublicUserData to use in
// cookie store.
func (u User) ToPublicUserData() *PublicUserData {
	return &PublicUserData{
		ID:        u.ID,
		Username:  u.Username,
		CreatedOn: u.CreatedOn,
		LastLogin: u.LastLogin,
	}
}

// PublicUserData represents user information
// to store in cookies for authorization.
type PublicUserData struct {
	ID        int
	Username  string
	CreatedOn time.Time
	LastLogin time.Time
}
