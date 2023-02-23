package main

import (
	"fmt"
	"bytes"
	"database/sql"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"strconv"
)

var (
	id       int
	title    string
	author   string
	date     string
	category string
	content  string
)

type AllThreads []Threads

type Threads struct {
	Id       string `json:"id,omitempty"`
	Title    string `json:"title,omitempty"`
	Author   string `json:"author,omitempty"`
	Date     string `json:"date,omitempty"`
	Category string `json:"category,omitempty"`
	Content  string `json:"content,omitempty"`
}

func GetAllThreads(w http.ResponseWriter, r *http.Request) {
	threadDb, _ := sql.Open("sqlite3", "./data.db")
	var AllThreads []Threads

	rows, err := threadDb.Query("SELECT id, title, author, date, category, content FROM threads")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(&id, &title, &author, &date, &category, &content)
		if err != nil {
			log.Fatal(err)
		}
		AllThreads = append(AllThreads, Threads{Id: strconv.Itoa(id), Title: title, Author: author, Date: date, Category: category, Content: content})
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}

	j, err := json.Marshal(AllThreads)
	if err != nil {
		http.Error(w, "couldn't retrieve threads", http.StatusInternalServerError)
	}
	
	w.Header().Set("Content-Type", "application/json")
	io.Copy(w, bytes.NewReader(j))
}
////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
type ThreadId struct {
	Id string `json:"threadId"`
}

func GetThread(w http.ResponseWriter, r *http.Request) {
	//gets data through header(POST; JSON)
	var PostId ThreadId

	err := json.NewDecoder(r.Body).Decode(&PostId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	fmt.Printf("%+v\n", PostId.Id)

	threadDb, _ := sql.Open("sqlite3", "./data.db")
	var AllThreads []Threads

	rows, err := threadDb.Query("SELECT id, title, author, date, category, content FROM threads WHERE id = ?", PostId.Id)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(&id, &title, &author, &date, &category, &content)
		if err != nil {
			log.Fatal(err)
		}
		AllThreads = append(AllThreads, Threads{Id: strconv.Itoa(id), Title: title, Author: author, Date: date, Category: category, Content: content})
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}

	j, err := json.Marshal(AllThreads)
	if err != nil {
		http.Error(w, "couldn't retrieve threads", http.StatusInternalServerError)
	}
	
	w.Header().Set("Content-Type", "application/json")
	io.Copy(w, bytes.NewReader(j))
}