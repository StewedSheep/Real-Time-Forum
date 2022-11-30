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

func GetRandomLaw(w http.ResponseWriter, r *http.Request) {
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
