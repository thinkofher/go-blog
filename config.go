package main

import (
	"github.com/thinkofher/go-blog/app/utils"
	"github.com/thinkofher/go-blog/db"
)

// SESSIONNAME represents default name of session for
// go-blog app.
const SESSIONNAME = "goblog-session"

// APPCONFIG represents configuration of application.
// Feel free to choose your own values.
var APPCONFIG = utils.AppConfig{
	SessionName:   "goblog-session",
	UserCookieKey: "user-cookie",
}

// DBCONFIG of the database.
// TODO: Make use of container and its envs in future.
var DBCONFIG = db.PSQLConfig{
	Host:     "localhost",
	Port:     5432,
	User:     "postgres",
	Password: "secret_postgres",
	DBName:   "goblog",
}
