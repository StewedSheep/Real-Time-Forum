package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
)

type ThreadData struct {
	Title    string `json:"title"`
	Category string `json:"category"`
	Content  string `json:"content"`
}

type ThreadResponse struct {
	thResp string
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

}
