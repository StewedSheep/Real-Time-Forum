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

func GetMessagesHandler(w http.ResponseWriter, r *http.Request) {
	keys, ok := r.URL.Query()["from"]
	if !ok || len(keys[0]) < 1 {
		log.Println("Url Param 'from' is missing")
		return
	}
	from := keys[0]

	keys, ok = r.URL.Query()["to"]
	if !ok || len(keys[0]) < 1 {
		log.Println("Url Param 'to' is missing")
		return
	}
	to := keys[0]

	database, _ := sql.Open("sqlite3", "./messages.db")
	msgs, err := GetMessages(database, from, to)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	log.Printf("getting message history", msgs)
	json.NewEncoder(w).Encode(msgs)
}
