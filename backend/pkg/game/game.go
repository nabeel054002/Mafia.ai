package game

import "github.com/google/uuid"

type GameState struct {
	Players   []string          `json:"players"`
	Votes     map[string]string `json:"votes"`
	Phase     string            `json:"phase"`
	Eliminated []string         `json:"eliminated"`
}

var games = make(map[string]*GameState)

func CreateGame() string {
	gameID := uuid.New().String()
	games[gameID] = &GameState{
		Players:   []string{},
		Votes:     make(map[string]string),
		Phase:     "waiting",
		Eliminated: []string{},
	}
	return gameID
}

func JoinGame(gameID string) string {
	playerID := uuid.New().String()
	game := games[gameID]
	game.Players = append(game.Players, playerID)
	return playerID
}

func GetGameState(gameID string) *GameState {
	return games[gameID]
}

func CastVote(gameID, playerID, voteTarget string) {
	game := games[gameID]
	if game.Phase == "voting" {
		game.Votes[playerID] = voteTarget
	}
}
