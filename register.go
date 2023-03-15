package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"net/mail"
	"time"
	"unicode"

	"golang.org/x/crypto/bcrypt"
)

type registerInf struct {
	Email     string `json:"email"`
	Username  string `json:"username"`
	Password  string `json:"password"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Gender    string `json:"gender"`
	Birthday  string `json:"birthday"`
}

type registerResponseData struct {
	RegisterError string
}

func ValidateRegister(w http.ResponseWriter, r *http.Request) {
	//gets data through header(POST; JSON)
	var regDat registerInf
	var regRespDat registerResponseData
	err := json.NewDecoder(r.Body).Decode(&regDat)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	fmt.Printf("%+v\n", regDat)

	userDb, err = sql.Open("sqlite3", "./user.db")
	if err != nil {
		panic(err.Error())
	}

	// check username for only alphaNumeric characters
	var nameAlphaNumeric = true
	for _, char := range regDat.Username {
		if !unicode.IsLetter(char) && !unicode.IsNumber(char) {
			nameAlphaNumeric = false
		}
	}
	// check Username length
	var nameLength bool
	if 5 <= len(regDat.Username) && len(regDat.Username) <= 16 {
		nameLength = true
	}
	// check password criteria
	fmt.Println("password:", regDat.Password, "\npswdLength:", len(regDat.Password))
	// variables that must pass for password creation criteria
	var pswdLowercase, pswdUppercase, pswdNumber, pswdLength, pswdNoSpaces bool
	pswdNoSpaces = true
	for _, char := range regDat.Password {
		switch {
		case unicode.IsLower(char):
			pswdLowercase = true
		case unicode.IsUpper(char):
			pswdUppercase = true
		case unicode.IsNumber(char):
			pswdNumber = true
		case unicode.IsSpace(int32(char)):
			pswdNoSpaces = false
		}
	}
	if 6 <= len(regDat.Password) && len(regDat.Password) <= 16 {
		pswdLength = true
	}
	emailFormat := false
	{
		_, err := mail.ParseAddress(regDat.Email)
		if err == nil {
			emailFormat = true
		}
	}
	fmt.Println("pswdLowercase:", pswdLowercase, "\npswdUppercase:", pswdUppercase, "\npswdNumber:", pswdNumber, "\npswdLength:", pswdLength, "\npswdNoSpaces:", pswdNoSpaces, "\nnameAlphaNumeric:", nameAlphaNumeric, "\nnameLength:", nameLength)
	if !pswdLowercase || !pswdUppercase || !pswdNumber || !pswdLength || !pswdNoSpaces || !nameAlphaNumeric || !nameLength || !emailFormat {
		regRespDat.RegisterError = "Please check the criteriums for registering."
		if err := json.NewEncoder(w).Encode(regRespDat.RegisterError); err != nil {
			panic(err)
		}
	}
	var uID string
	{
		// check if username already exists in db
		stmt := "SELECT userID FROM bcrypt WHERE username = ?"
		row := userDb.QueryRow(stmt, regDat.Username)
		err := row.Scan(&uID)
		if err != sql.ErrNoRows {
			regRespDat.RegisterError = "Username already exists."
			if err := json.NewEncoder(w).Encode(regRespDat.RegisterError); err != nil {
				panic(err)
			}
		}
	}
	{
		// check if email already exists in db
		stmt := "SELECT userID FROM bcrypt WHERE email = ?"
		row := userDb.QueryRow(stmt, regDat.Email)
		err := row.Scan(&uID)
		if err != sql.ErrNoRows {
			fmt.Println(err)
			regRespDat.RegisterError = "Email is already taken."
			if err := json.NewEncoder(w).Encode(regRespDat.RegisterError); err != nil {
				panic(err)
			}
		}
	}
	{
		//AgeCheck
		t := time.Now()
		t1, _ := time.Parse("2 Jan 06 03:04PM", regDat.Birthday)
		t2 := t.AddDate(-13, 0, 0)
		if t2.Before(t1) {
			regRespDat.RegisterError = "Age is less than 13y"
			if err := json.NewEncoder(w).Encode(regRespDat.RegisterError); err != nil {
				panic(err)
			}
		}
	}
	{
		//Name Strict Letters check
		namesLetters := true
		for _, char := range regDat.FirstName {
			if !unicode.IsLetter(char) {
				namesLetters = false
			}
		}
		for _, char := range regDat.LastName {
			if !unicode.IsLetter(char) {
				namesLetters = false
			}
		}
		if !namesLetters {
			regRespDat.RegisterError = "First and last name only in latin letters."
			if err := json.NewEncoder(w).Encode(regRespDat.RegisterError); err != nil {
				panic(err)
			}
		}
	}
	{
		//Gender check
		if regDat.Gender == ""{
			regRespDat.RegisterError = "Gender field is empty."
			if err := json.NewEncoder(w).Encode(regRespDat.RegisterError); err != nil {
				panic(err)
			}
		}
	}
	if regRespDat.RegisterError == "" {
		// create hash from password
		var hash []byte
		hash, err = bcrypt.GenerateFromPassword([]byte(regDat.Password), bcrypt.DefaultCost)
		if err != nil {
			regRespDat.RegisterError = "there was a problem registering account(password encrypting is broken)"
			if err := json.NewEncoder(w).Encode(regRespDat.RegisterError); err != nil {
				panic(err)
			}
		}
		fmt.Println("crypted pswd (hash):", string(hash))
		//inserts user into users database
		var insertStmt *sql.Stmt
		insertStmt, err = userDb.Prepare("INSERT INTO bcrypt (username, hash, email, age, gender, firstName, lastName) VALUES (?, ?, ?, ?, ?, ?, ?);")
		if err != nil {
			regRespDat.RegisterError = "there was a problem registering account(cant insert data into database)"
			if err := json.NewEncoder(w).Encode(regRespDat.RegisterError); err != nil {
				panic(err)
			}
		}
		defer insertStmt.Close()
		var result sql.Result
		result, err := insertStmt.Exec(regDat.Username, hash, regDat.Email, regDat.Birthday, regDat.Gender, regDat.FirstName, regDat.LastName)
		rowsAff, _ := result.RowsAffected()
		lastIns, _ := result.LastInsertId()
		fmt.Println("rowsAff:", rowsAff)
		fmt.Println("lastIns:", lastIns)
		if err != nil {
			regRespDat.RegisterError = "error inserting new user into user db."
			if err := json.NewEncoder(w).Encode(regRespDat.RegisterError); err != nil {
				panic(err)
			}
		}
		regRespDat.RegisterError = "Congratulations! you have successfully registered your account."
		if err := json.NewEncoder(w).Encode(regRespDat.RegisterError); err != nil {
			panic(err)
		}
	}
}
