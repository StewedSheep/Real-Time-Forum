package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	uuid "github.com/satori/go.uuid"
	"golang.org/x/crypto/bcrypt"
)

type loginInf struct {
	Username string `json:"userName"`
	Password string `json:"password"`
}

type loginResponseData struct {
	LoginError string
}

func LoginAuth(w http.ResponseWriter, r *http.Request) {
	//gets data through header(POST; JSON)
	var logDat loginInf
	var logRespDat loginResponseData

	err := json.NewDecoder(r.Body).Decode(&logDat)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	fmt.Printf("%+v\n", logDat)

	userDb, err = sql.Open("sqlite3", "./user.db")
	if err != nil {
		panic(err.Error())
	}

	var hash string
	var username string

	stmt := "SELECT hash,username FROM bcrypt WHERE username = ? OR email = ?"
	row := userDb.QueryRow(stmt, logDat.Username, logDat.Username)
	err = row.Scan(&hash, &username)
	fmt.Println(username)
	if err != nil {
		fmt.Println(err)
		fmt.Println("error selecting Hash in db by Username")
	}
	err = bcrypt.CompareHashAndPassword([]byte(hash), []byte(logDat.Password))
	if err == nil {
		fmt.Println("Login Successful")

		cookie := http.Cookie{
			Name:    "session_token",
			Path:    "/",
			Value:   "",
			Expires: time.Now().Add(-1 * time.Second),
		}
		http.SetCookie(w, &cookie)
		//creates new cookie
		sessionToken := uuid.NewV4().String() + "::" + logDat.Username
		expiresAt := time.Now().Add(36000 * time.Second)

		//writes cookie to header
		http.SetCookie(w, &http.Cookie{
			Path:    "/",
			Name:    "session_token",
			Value:   sessionToken,
			Expires: expiresAt,
		})
		// sends to index page

		return
	}
	logRespDat.LoginError = "Check login info."
	if err := json.NewEncoder(w).Encode(logRespDat.LoginError); err != nil {
		panic(err)
	}
}

func ResetCookie(w http.ResponseWriter, r *http.Request) {
	cookie := http.Cookie{
		Name:    "session_token",
		Path:    "/",
		Value:   "",
		MaxAge:  -1,
		Expires: time.Now().Add(-100 * time.Hour),
	}
	fmt.Println("deleted cookies")
	http.SetCookie(w, &cookie)
}
