package main

import (
	"net/http"

	"github.com/SuperTikuwa/matching_game/handler"
)

func init() {

}

func main() {
	http.HandleFunc("/words", handler.WordsGetHandler)
	http.ListenAndServe(":80", nil)
}
