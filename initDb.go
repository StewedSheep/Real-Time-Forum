package main

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

var userDb *sql.DB
var threadDb *sql.DB
var chatDb *sql.DB

func StartDb() {
	var err error
	userDb, err = sql.Open("sqlite3", "./user.db")
	if err != nil {
		panic(err.Error())
	}
	statement, _ := userDb.Prepare(`CREATE TABLE IF NOT EXISTS bcrypt
	(
	userID INTEGER PRIMARY KEY,
	username TEXT,
	hash TEXT,
	email TEXT,
	age TEXT,
	gender TEXT,
	firstName TEXT,
	lastName TEXT
	)`)
	statement.Exec()
	defer userDb.Close()

	threadDb, err = sql.Open("sqlite3", "./data.db")
	if err != nil {
		panic(err.Error())
	}
	statement, _ = threadDb.Prepare("CREATE TABLE IF NOT EXISTS threads (id INTEGER PRIMARY KEY, title TEXT, author TEXT, date TEXT,category TEXT, content Text)")
	statement.Exec()
	statement, _ = threadDb.Prepare("CREATE TABLE IF NOT EXISTS comments (id INTEGER PRIMARY KEY, relId INTEGER, author TEXT, date TEXT, content Text)")
	statement.Exec()
	defer threadDb.Close()

	chatDb, err := sql.Open("sqlite3", "./messages.db")
	if err != nil {
		panic(err.Error())
	}
	statement, err = chatDb.Prepare("CREATE TABLE IF NOT EXISTS messages (id INTEGER PRIMARY KEY, author TEXT, recipient TEXT, content TEXT, date DATETIME, read INTEGER DEFAULT 0)")
	if err != nil {
		panic(err.Error())
	}
	statement.Exec()
	defer chatDb.Close()
}
