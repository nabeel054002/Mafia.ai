package routes

import (
	"encoding/json"
	"net/http"
	"github.com/gorilla/handlers"
	"github.com/gorilla/websocket"
	"mafia.ai-backend/pkg/game"
)

func InitializeRoutes() {
	// Wrap routes with CORS handler
	http.HandleFunc("/join", JoinOrCreateGame)
	http.HandleFunc("/status", FetchGameState)
	http.HandleFunc("/vote", CastVote)
	http.HandleFunc("/health", HealthCheck)
	http.Handle("/", handlers.CORS(handlers.AllowedOrigins([]string{"*"}))(http.DefaultServeMux))
}

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool { return true },
}

func FetchGameState(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	gameID := r.URL.Query().Get("gameID")
	gameState := game.GetGameState(gameID)
	json.NewEncoder(w).Encode(gameState)
}

func CastVote(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	gameID := r.URL.Query().Get("gameID")
	playerID := r.URL.Query().Get("playerID")
	voteTarget := r.URL.Query().Get("voteTarget")
	game.CastVote(gameID, playerID, voteTarget)
	w.Write([]byte("Vote casted"))
}

// HealthCheck returns a simple status message to ensure the server is running
func HealthCheck(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	w.Write([]byte("OK"))
}

// API Route to create or join a game
func JoinOrCreateGame(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)

	// Check if game instance exists
	gameID := r.URL.Query().Get("gameID")
	var gameState game.GameState
	if gameID != "" {
		// Fetch game state to see if game already exists
		// here game refers to the pkg 
		gameState = *game.GetGameState(gameID)
	}

	// If game exists and has less than 5 players, join the game
	if gameID != "" && len(gameState.Players) < 5 {
		playerID := game.JoinGame(gameID)
		json.NewEncoder(w).Encode(map[string]string{"gameID": gameID, "playerID": playerID})
	} else if gameID == "" || len(gameState.Players) >= 5 {
		// If game doesn't exist or is full, create a new game
		newGameID := game.CreateGame()

		// Create players and start the game
		playerID := game.JoinGame(newGameID)
		if len(gameState.Players) == 5 {
			// Create AI players once 5 players are reached
			game.AddAIPlayers(newGameID)
			game.StartGame(newGameID)
		}

		// Respond with new game ID and player ID
		json.NewEncoder(w).Encode(map[string]string{"gameID": newGameID, "playerID": playerID})
	}
}

func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
}

