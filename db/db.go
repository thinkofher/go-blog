package db

import (
	"database/sql"
	"errors"
	"fmt"

	_ "github.com/lib/pq"
)

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

// DBWrapper wraps PSQL database and provide friendly API
// to handle common operatons without writing additional
// SQL code.
type DBWrapper struct {
	DB *sql.DB
}

// NewDBWrapper returns wrapped and opened database.
// Performs single ping to that database. Returns error
// whenever something goes wrong.
func NewDBWrapper(config PSQLConfig) (DBWrapper, error) {
	db, err := sql.Open("postgres", config.String())
	if err != nil {
		return DBWrapper{}, err
	}

	err = db.Ping()
	if err != nil {
		return DBWrapper{}, err
	}

	return DBWrapper{DB: db}, nil
}

// SetUser registers given User model in wrapped database.
// TODO: Test it.
func (wrapper DBWrapper) SetUser(user User) error {
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

func (wrapper DBWrapper) GetUser(username string) (User, error) {
	user := User{}
	statement := `
	SELECT
		user_id, username, password,
		email, created_on, last_login
	FROM
		blog_user
	WHERE
		username = $1;
	`
	row := wrapper.DB.QueryRow(statement, username)

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
