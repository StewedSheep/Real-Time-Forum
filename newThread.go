package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
)

type ThreadData struct {
	Title    string `json:"title"`
	Author   string `json:"author"`
	Date     string `json:"date"`
	Category string `json:"category"`
	Content  string `json:"content"`
}

func CreateThread(w http.ResponseWriter, r *http.Request) {
	var thDat ThreadData
	//var thRespDat ThreadResponse
	err := json.NewDecoder(r.Body).Decode(&thDat)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	fmt.Printf("%+v\n", thDat)

	threadDb, err = sql.Open("sqlite3", "./data.db")
	if err != nil {
		panic(err.Error())
	}

	var insertStmt *sql.Stmt
	insertStmt, err = threadDb.Prepare("INSERT INTO threads (title, author, date, category, content) VALUES (?, ?, ?, ?, ?);")
	if err != nil {
		fmt.Println("there was a problem registering account(cant insert data into database)")
	}

	var result sql.Result
	fmt.Println(thDat.Title, thDat.Author, thDat.Date, thDat.Category, thDat.Content)
	result, err = insertStmt.Exec(thDat.Title, thDat.Author, thDat.Date, thDat.Category, thDat.Content)
	rowsAff, _ := result.RowsAffected()
	lastIns, _ := result.LastInsertId()
	fmt.Println("rowsAff:", rowsAff)
	fmt.Println("lastIns:", lastIns)
	if err != nil {
		fmt.Println("success registering.")
	}
}
