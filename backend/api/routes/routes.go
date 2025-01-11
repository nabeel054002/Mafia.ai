package routes

import (
	"encoding/json"
	"net/http"
	"backend/pkg/game"
	"backend/internal/utils"
)

func InitializeRoutes() {
	http.HandleFunc("/create", CreateGame)
	http.HandleFunc("/join", JoinGame)
	http.HandleFunc("/status", FetchGameState)
	http.HandleFunc("/vote", CastVote)
}

func CreateGame(w http.ResponseWriter, r *http.Request) {
	gameID := game.CreateGame()
	json.NewEncoder(w).Encode(map[string]string{"gameID": gameID})
}

func JoinGame(w http.ResponseWriter, r *http.Request) {
	gameID := r.URL.Query().Get("gameID")
	playerID := game.JoinGame(gameID)
	json.NewEncoder(w).Encode(map[string]string{"playerID": playerID})
}

func FetchGameState(w http.ResponseWriter, r *http.Request) {
	gameID := r.URL.Query().Get("gameID")
	gameState := game.GetGameState(gameID)
	json.NewEncoder(w).Encode(gameState)
}

func CastVote(w http.ResponseWriter, r *http.Request) {
	gameID := r.URL.Query().Get("gameID")
	playerID := r.URL.Query().Get("playerID")
	voteTarget := r.URL.Query().Get("voteTarget")
	game.CastVote(gameID, playerID, voteTarget)
	w.Write([]byte("Vote casted"))
}
