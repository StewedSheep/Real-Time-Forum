package main

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"strconv"
)

type DbMessage struct {
	Id      int    `json:"id"`
	From    string `json:"from"`
	To      string `json:"to"`
	Content string `json:"content"`
	Date    string `json:"date"`
	Read    bool   `json:"read"`
}

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

//handles the request from frontend to get messages for chat
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

	keys, ok = r.URL.Query()["page"]
	if !ok || len(keys[0]) < 1 {
		log.Println("Url Param 'page' is missing")
		return
	}
	page, _ := strconv.Atoi(keys[0])

	database, _ := sql.Open("sqlite3", "./messages.db")
	msgs, err := GetMessages(database, from, to, page)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	log.Printf("getting message history", msgs)
	json.NewEncoder(w).Encode(msgs)
}

// //
// func markMessageAsRead(db *sql.DB, from string, to string) error {
// 	statement, err := db.Prepare("UPDATE messages SET read = 1 WHERE author = ?")
// 	if err != nil {
// 		return err
// 	}
// 	_, err = statement.Exec(id)
// 	return err
// }

func GetMessages(db *sql.DB, from string, to string, page int) ([]DbMessage, error) {
	offset := (page - 1) * 10
	rows, err := db.Query("SELECT id, author, recipient, content, date FROM messages WHERE author = ? AND recipient = ? OR author = ? AND recipient = ? ORDER BY id DESC LIMIT 10 OFFSET ?", from, to, to, from, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var msgs []DbMessage
	for rows.Next() {
		var msg DbMessage
		err := rows.Scan(&msg.Id, &msg.From, &msg.To, &msg.Content, &msg.Date)
		if err != nil {
			return nil, err
		}
		msgs = append([]DbMessage{msg}, msgs...)
	}

	err = rows.Err()
	if err != nil {
		return nil, err
	}

	return msgs, nil
}
