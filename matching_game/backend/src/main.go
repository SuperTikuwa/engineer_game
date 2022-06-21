package main

import (
	"net/http"

	"github.com/SuperTikuwa/matching_game/handler"
	"github.com/gorilla/mux"
)

func init() {

}

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/words", handler.WordsGetHandler).Methods("GET")
	r.HandleFunc("/game", handler.GamePostHandler).Methods("POST")
	r.HandleFunc("/game/{game_id}", handler.GamePutHandler).Methods("PUT")
	http.ListenAndServe(":80", r)
}
