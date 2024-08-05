package main

import (
	"log"
	"net/http"
	"sync"

	"github.com/gorilla/websocket"
)

var clients = make(map[*websocket.Conn]string) // connected clients
var broadcast = make(chan Message)             // broadcast channel
var mu sync.Mutex

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

// Message defines the structure of a chat message
type Message struct {
	Type     string   `json:"type"`
	Username string   `json:"username"`
	Message  string   `json:"message"`
	Users    []string `json:"users"`
}

func main() {
	fs := http.FileServer(http.Dir("./public"))
	http.Handle("/", fs)

	http.HandleFunc("/ws", handleConnections)

	go handleMessages()

	log.Println("HTTP server started on :8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

func handleConnections(w http.ResponseWriter, r *http.Request) {
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer ws.Close()

	for {
		var msg Message
		err := ws.ReadJSON(&msg)
		if err != nil {
			log.Printf("error: %v", err)
			mu.Lock()
			delete(clients, ws)
			mu.Unlock()
			updateUsersList()
			break
		}
		handleMessage(ws, msg)
	}
}

func handleMessage(ws *websocket.Conn, msg Message) {
	switch msg.Type {
	case "login":
		mu.Lock()
		clients[ws] = msg.Username
		mu.Unlock()
		updateUsersList()
	case "message":
		broadcast <- msg
	}
}

func updateUsersList() {
	usernames := make([]string, 0, len(clients))
	for _, username := range clients {
		usernames = append(usernames, username)
	}
	userListMessage := Message{
		Type:    "users",
		Message: "",
		Users:   usernames,
	}
	for client := range clients {
		client.WriteJSON(userListMessage)
	}
}

func handleMessages() {
	for {
		msg := <-broadcast
		for client := range clients {
			err := client.WriteJSON(msg)
			if err != nil {
				log.Printf("error: %v", err)
				client.Close()
				mu.Lock()
				delete(clients, client)
				mu.Unlock()
				updateUsersList()
			}
		}
	}
}
