package main

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"io"
	"log"
	"net/http"
)
var username string

func GetAllUsers(w http.ResponseWriter, r *http.Request) {
	userDb, _ := sql.Open("sqlite3", "./user.db")
	var AllUsers []string

	rows, err := userDb.Query("SELECT username FROM bcrypt")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(&username)
		if err != nil {
			log.Fatal(err)
		}
		AllUsers = append(AllUsers, username)
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}

	j, err := json.Marshal(AllUsers)
	if err != nil {
		http.Error(w, "couldn't retrieve users", http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	io.Copy(w, bytes.NewReader(j))
}