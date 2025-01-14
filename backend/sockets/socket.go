package sockets

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func StartSocketServer() {
	http.HandleFunc("/ws", handleConnections)
	log.Println("WebSocket server started on /ws")
}

func handleConnections(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Printf("WebSocket upgrade failed: %v\n", err)
		return
	}
	defer conn.Close()

	// Handle incoming messages
	for {
		messageType, message, err := conn.ReadMessage()
		if err != nil {
			log.Printf("Error reading message: %v\n", err)
			break
		}
		log.Printf("Received: %s\n", message)

		// Echo message back to client
		err = conn.WriteMessage(messageType, message)
		if err != nil {
			log.Printf("Error writing message: %v\n", err)
			break
		}
	}
}
