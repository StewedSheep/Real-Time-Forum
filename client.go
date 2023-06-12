package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

type Client struct {
	conn     *websocket.Conn
	outbound chan Message
	username string
}

type Message struct {
	To      string `json:"to"`
	From    string `json:"from"`
	Content string `json:"content"`
	Date    string `json:"date"`
	Read    bool
}

var clients = make(map[string]*Client)
var register = make(chan *Client)
var messages = make(chan Message)
var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func Echo() {
	for {
		select {
		case message := <-messages:
			if clients[message.To] != nil {
				clients[message.To].outbound <- message
			}
		case client := <-register:
			fmt.Println("registered ws client", client.username)
			clients[client.username] = client
		}
	}
}

func WsHandler(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}

	username := r.URL.Query().Get("username")
	client := &Client{
		conn:     conn,
		outbound: make(chan Message),
		username: username,
	}

	go client.read()
	go client.write()

	register <- client
}

func (c *Client) read() {
	defer c.conn.Close()
	for {
		msg := Message{}
		if err := c.conn.ReadJSON(&msg); err != nil {
			log.Println(err)
			return
		}
		msg.From = c.username
		log.Printf("incoming message %+v", msg)

		database, _ := sql.Open("sqlite3", "./messages.db")
		err := insertMessage(database, msg.From, msg.To, msg.Content, msg.Date)
		if err != nil {
			log.Printf("error inserting message into database: %v", err)
		}

		messages <- msg
	}
}

//sends message if both users are online
func (c *Client) write() {
	defer c.conn.Close()
	for msg := range c.outbound {
		log.Printf("Sending message %+v", msg)
		if err := c.conn.WriteJSON(msg); err != nil {
			log.Println(err)
			return
		}
	}
}

func insertMessage(db *sql.DB, from string, to string, content string, date string) error {
	statement, err := db.Prepare("INSERT INTO messages (author, recipient, content, date) VALUES (?, ?, ?, ?)")
	if err != nil {
		return err
	}
	_, err = statement.Exec(from, to, content, date)
	return err
}

// move to chatdb.go when done debugging
