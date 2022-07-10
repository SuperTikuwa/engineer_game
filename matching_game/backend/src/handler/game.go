package handler

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/SuperTikuwa/matching_game/redisoperation"
	"github.com/gorilla/mux"
)

func GameGetHandler(w http.ResponseWriter, r *http.Request) {
	g := redisoperation.GetGames()

	games := struct {
		Games []string `json:"games"`
	}{
		g,
	}

	json, err := json.Marshal(games)
	if err != nil {
		log.Println(err)
	}

	fmt.Fprint(w, string(json))
}

func GamePostHandler(w http.ResponseWriter, r *http.Request) {
	gameID := redisoperation.GenerateGameID()

	j := struct {
		ID string `json:"gameID"`
	}{
		gameID,
	}

	res, err := json.Marshal(j)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Fprintln(w, string(res))
}

func GamePutHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	redisoperation.JoinGame(vars["game_id"], r.URL.Query().Get("player_id"))
}
