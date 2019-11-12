package main

import (
	"github.com/thinkofher/go-blog/db"
)

var CONFIG = db.PSQLConfig{
	Host:     "localhost",
	Port:     5432,
	User:     "postgres",
	Password: "secret_postgres",
	DBName:   "goblog",
}
