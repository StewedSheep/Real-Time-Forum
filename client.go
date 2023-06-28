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
	Type    string      `json:type`
	To      string      `json:"to"`
	From    string      `json:"from"`
	Content interface{} `json:"content"`
	Date    string      `json:"date"`
}

type User struct {
	ID                int
	Username          string
	LatestMessageDate string
}

var clients = make(map[string]*Client)
var register = make(chan *Client)
var unregister = make(chan *websocket.Conn)
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
			sendActiveUsers()
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

	sendActiveUsers()

	// Listen for close events from the client
	go func() {
		for {
			_, _, err := conn.ReadMessage()
			if err != nil {
				if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
					log.Printf("Websocket error: %v", err)
				}
				break
			}
		}
		// Unregister the client connection
		fmt.Println("deleted ws conn: ", client.username)
		delete(clients, client.username)
		sendActiveUsers()
	}()
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
		err := insertMessage(database, msg.From, msg.To, msg.Content.(string), msg.Date)
		if err != nil {
			log.Printf("error inserting message into database: %v", err)
		}

		messages <- msg
		sendUserMessages(c)
	}
}

//sends message if both users are online
func (c *Client) write() {
	defer c.conn.Close()
	for msg := range c.outbound {
		log.Printf("Sending message %+v to %v", msg, c.username)
		if err := c.conn.WriteJSON(msg); err != nil {
			log.Println(err)
			return
		}
	}
}

//checks for closed connections via ws
// func isConnectionClosed(conn *websocket.Conn) bool {
// 	_, _, err := conn.ReadMessage()
// 	if err != nil {
// 		if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
// 			log.Printf("Websocket error: %v", err)
// 		}
// 		return true
// 	}
// 	return false
// }

//adds message to messages.db
func insertMessage(db *sql.DB, from string, to string, content string, date string) error {
	statement, err := db.Prepare("INSERT INTO messages (author, recipient, content, date) VALUES (?, ?, ?, ?)")
	if err != nil {
		return err
	}
	_, err = statement.Exec(from, to, content, date)
	return err
}

//broadcasts active users to all clients
func sendActiveUsers() {
	activeUsers := make([]string, 0, len(clients))
	for username := range clients {
		activeUsers = append(activeUsers, username)
	}

	for _, client := range clients {
		client.outbound <- Message{
			Type:    "activeUsers",
			To:      "",
			From:    "",
			Content: activeUsers,
			Date:    "",
		}
	}
}

//broadcasts latest messages from users to sort them in chatbar
func sendUserMessages(c *Client) {
	// Connect to user.db
	userDB, err := sql.Open("sqlite3", "./user.db")
	if err != nil {
		log.Fatal(err)
	}
	defer userDB.Close()

	// Query all users from user.db
	userRows, err := userDB.Query("SELECT userID, username FROM bcrypt")
	if err != nil {
		log.Fatal(err)
	}
	defer userRows.Close()

	// Prepare a map to store the latest message date for each user
	userLatestMessageDates := make(map[int]string)

	// Loop through each user
	for userRows.Next() {
		var user User
		err := userRows.Scan(&user.ID, &user.Username)
		if err != nil {
			log.Fatal(err)
		}

		// Connect to messages.db
		messagesDB, err := sql.Open("sqlite3", "./messages.db")
		if err != nil {
			log.Fatal(err)
		}
		defer messagesDB.Close()

		// Query the latest message date for the current user from messages.db
		var latestMessageDate sql.NullString
		err = messagesDB.QueryRow("SELECT MAX(date) FROM messages WHERE author = ? OR recipient = ?", user.ID, user.ID).Scan(&latestMessageDate)
		if err != nil {
			if err == sql.ErrNoRows {
				// Handle case where user has no messages
				latestMessageDate.String = "No messages"
			} else {
				log.Fatal(err)
			}
		}

		// Check if the latest message date is empty or null, provide a default value if so
		if !latestMessageDate.Valid || latestMessageDate.String == "" {
			latestMessageDate.String = "No date available"
		}

		// Store the latest message date in the map
		userLatestMessageDates[user.ID] = latestMessageDate.String
	}
	c.outbound <- Message{
		Type:    "listMsgs",
		To:      "",
		From:    "",
		Content: userLatestMessageDates,
		Date:    "",
	}
}
