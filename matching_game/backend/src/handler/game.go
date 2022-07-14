package handler

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/SuperTikuwa/matching_game/redisoperation"
)

func GamePostHandler(w http.ResponseWriter, r *http.Request) {
	gameID := redisoperation.FindEmptyGame()
	if gameID == "" {
		gameID = redisoperation.GenerateGameID()
		redisoperation.CreateGame(gameID, "12345")

		w.WriteHeader(http.StatusCreated)
		fmt.Fprint(w, `{"gameID":"`+gameID+`"}`)
		return
	}

	gameID = strings.Split(gameID, ":")[1]

	redisoperation.JoinGame(gameID, "54321")
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, `{"gameID":"`+gameID+`"}`)
}
