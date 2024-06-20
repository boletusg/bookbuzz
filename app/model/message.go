package model

import (
	"fmt"
	"github.com/gorilla/websocket"
	_ "github.com/gorilla/websocket"
	"net/http"
)

func MessengerHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		// Отображение страницы login.html
		http.ServeFile(w, r, "public/html/messenger_page.html")
	}
}

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func HandleWebSocket(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println("Error upgrading to WebSocket:", err)
		return
	}
	defer conn.Close()

	for {
		messageType, message, err := conn.ReadMessage()
		if err != nil {
			fmt.Println("Error reading message:", err)
			return
		}

		fmt.Printf("Received message: %s\n", message)
		if err := conn.WriteMessage(messageType, message); err != nil {
			fmt.Println("Error writing message:", err)
			return
		}
	}
}
