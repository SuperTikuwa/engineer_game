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

	r.HandleFunc("/words", handler.WordsGetHandler).Methods("GET").Queries("mode", "{mode:[0-2]}", "gameID", "{gameID}")

	r.HandleFunc("/games", handler.GamePostHandler).Methods("POST")
	http.ListenAndServe(":80", r)
}
