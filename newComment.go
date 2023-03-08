package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
)

type CommentData struct {
	RelId   string `json:"relId"`
	Author  string `json:"author"`
	Date    string `json:"date"`
	Content string `json:"content"`
}

func CreateComment(w http.ResponseWriter, r *http.Request) {
	var commDat CommentData
	//var thRespDat ThreadResponse
	err := json.NewDecoder(r.Body).Decode(&commDat)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	fmt.Println("Creating comment")
	fmt.Printf("%+v\n", commDat)

	threadDb, err = sql.Open("sqlite3", "./data.db")
	if err != nil {
		panic(err.Error())
	}

	var insertStmt *sql.Stmt
	insertStmt, err = threadDb.Prepare("INSERT INTO comments (RelId, author, date, content) VALUES (?, ?, ?, ?);")
	if err != nil {
		fmt.Println("there was a problem registering account(cant insert data into database)")
	}

	var result sql.Result
	fmt.Println(commDat.RelId, commDat.Author, commDat.Date, commDat.Content)
	result, err = insertStmt.Exec(commDat.RelId, commDat.Author, commDat.Date, commDat.Content)
	rowsAff, _ := result.RowsAffected()
	lastIns, _ := result.LastInsertId()
	fmt.Println("rowsAff:", rowsAff)
	fmt.Println("lastIns:", lastIns)
	if err != nil {
		fmt.Println("success registering.")
	}
}
