package handler

import (
	"net/http"

	"github.com/SuperTikuwa/matching_game/dboperation"
)

func WordsGetHandler(w http.ResponseWriter, r *http.Request) {
	dboperation.SelectWords(dboperation.STUDENT)
}
