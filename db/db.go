package db

import (
	"database/sql"
	"errors"
	"fmt"
)

// ErrNoUser is returned when there is no user with specific
// data in database.
var ErrNoUser = errors.New("db: no such user in database")

// PSQLConfig represents Postgres database config.
type PSQLConfig struct {
	Host     string
	Port     int
	User     string
	Password string
	DBName   string
}

// String method of PSQLConfig struct returns
// ready to use config string to open database.
// TODO: Test it.
func (c PSQLConfig) String() string {
	return fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		c.Host, c.Port, c.User, c.Password, c.DBName)
}

// Wrapper wraps PSQL database and provide friendly API
// to handle common operatons without writing additional
// SQL code.
type Wrapper struct {
	DB *sql.DB
}

// NewWrapper returns wrapped and opened database.
// Performs single ping to that database. Returns error
// whenever something goes wrong.
func NewWrapper(config PSQLConfig) (Wrapper, error) {
	db, err := sql.Open("postgres", config.String())
	if err != nil {
		return Wrapper{}, err
	}

	err = db.Ping()
	if err != nil {
		return Wrapper{}, err
	}

	return Wrapper{DB: db}, nil
}

// SetUser registers given User model in wrapped database.
// TODO: Test it.
func (wrapper Wrapper) SetUser(user User) error {
	statement := `
	INSERT INTO blog_user (username, password, email, created_on, last_login)
	VALUES
		($1, $2, $3, $4, $5);
	`
	_, err := wrapper.DB.Exec(
		statement, user.Username, user.HashedPassword,
		user.Email, user.CreatedOn, user.LastLogin)

	if err != nil {
		return err
	}

	return nil
}

func (wrapper Wrapper) queryUser(queryFunc func() *sql.Row) (User, error) {
	var user User

	row := queryFunc()

	err := row.Scan(&user.ID, &user.Username, &user.HashedPassword,
		&user.Email, &user.CreatedOn, &user.LastLogin)
	if err == sql.ErrNoRows {
		// This means, there is no such user in database
		return user, ErrNoUser
	} else if err != nil {
		panic(err)
	}

	return user, nil
}

// GetUser returns user, with given username, data in form of
// User struct.
// Returns ErrNoUser, when there are not any user with
// given username.
// Pancics, when there aren't appropriate table in database
// (check init.sql script for further information about data
// structures in database).
// TODO: Test it.
func (wrapper Wrapper) GetUser(username string) (User, error) {
	return wrapper.queryUser(func() *sql.Row {
		statement := `
		SELECT
			user_id, username, password,
			email, created_on, last_login
		FROM
			blog_user
		WHERE
			username = $1;
		`
		return wrapper.DB.QueryRow(statement, username)
	})
}

// GetUserByID returns user, with given ID, data in form of
// User struct.
// Returns ErrNoUser, when there are not any user with
// given username.
// Pancics, when there aren't appropriate table in database
// (check init.sql script for further information about data
// structures in database).
// TODO: Test it.
func (wrapper Wrapper) GetUserByID(id int) (User, error) {
	return wrapper.queryUser(func() *sql.Row {
		statement := `
		SELECT
			user_id, username, password,
			email, created_on, last_login
		FROM
			blog_user
		WHERE
			user_id = $1;
		`
		return wrapper.DB.QueryRow(statement, id)
	})
}

// SetPost inserts given Post struct into wrapped database.
// Return nil, when transaction ended with success.
func (wrapper Wrapper) SetPost(post Post) error {
	statement := `
	INSERT INTO post (author_id, body, created_on)
	VALUES
		($1, $2, $3);
	`
	_, err := wrapper.DB.Exec(
		statement, post.Author.ID, post.Body, post.CreatedOn)

	if err != nil {
		return err
	}

	return nil
}

// GetPosts returns all posts from database.
func (wrapper Wrapper) GetPosts() ([]Post, error) {
	statement := `
	SELECT
		author_id, body, created_on
	FROM
		post
	ORDER BY
		created_on DESC;
	`

	rows, err := wrapper.DB.Query(statement)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var posts []Post
	var post Post
	var userID int

	for rows.Next() {
		err = rows.Scan(&userID, &post.Body, &post.CreatedOn)
		if err != nil {
			return nil, err
		}

		user, err := wrapper.GetUserByID(userID)
		if err != nil {
			return nil, err
		}

		post.Author = user
		posts = append(posts, post)
	}

	err = rows.Err()
	if err != nil {
		return nil, err
	}

	return posts, nil
}
