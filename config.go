package main

import (
	"github.com/thinkofher/go-blog/db"
)

// CONFIG of the database.
// TODO: Make use of container and its envs in future.
var CONFIG = db.PSQLConfig{
	Host:     "localhost",
	Port:     5432,
	User:     "postgres",
	Password: "secret_postgres",
	DBName:   "goblog",
}
