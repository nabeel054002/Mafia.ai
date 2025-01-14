package main

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"
	"mafia.ai-backend/api/routes"
	"mafia.ai-backend/pkg/game"
)

func main() {
	// Start the HTTP server (handles /join, /status, /vote, etc.)
	go startHTTPServer()

	// Start the WebSocket handler in a separate goroutine
	go startWebSocketServer()

	log.Println("Server running on http://localhost:3000")
	select {} // Keep the main goroutine alive indefinitely
}

func startHTTPServer() {
	// Initialize and start the HTTP server (using existing routes)
	routes.InitializeRoutes()

	// Start the HTTP server
	log.Fatal(http.ListenAndServe(":3000", nil))
}

// WebSocket server to handle WebSocket connections
func startWebSocketServer() {
	// Example WebSocket server, can be customized for your use case
	http.HandleFunc("/socket", SocketHandler)
	log.Println("WebSocket server running on ws://localhost:3000/socket")
	log.Fatal(http.ListenAndServe(":3001", nil)) // Running on a different port (3001)
}

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool { return true },
}

// Example WebSocket handler
func SocketHandler(w http.ResponseWriter, r *http.Request) {
	// Upgrade HTTP connection to WebSocket
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		http.Error(w, "Could not open websocket connection", http.StatusBadRequest)
		return
	}
	defer conn.Close()

	// WebSocket logic (e.g., broadcasting game state)
	for {
		// Read message from WebSocket
		_, msg, err := conn.ReadMessage()
		if err != nil {
			log.Println("Error reading message:", err)
			break
		}

		// Handle the message and send a response
		gameID := string(msg) // Replace this logic with your game-related logic
		gameState := game.GetGameState(gameID)
		err = conn.WriteJSON(gameState)
		if err != nil {
			log.Println("Error writing message:", err)
			break
		}
	}
}
