package handler

import (
	"fmt"
	"log"
	"net/http"

	"github.com/SuperTikuwa/matching_game/dboperation"
	"github.com/SuperTikuwa/matching_game/models"
)

func WordsGetHandler(w http.ResponseWriter, r *http.Request) {
	m := r.URL.Query().Get("mode")
	gameID := r.URL.Query().Get("gameID")
	mode := ""
	switch m {
	case "0":
		mode = dboperation.STUDENT
	case "1":
		mode = dboperation.NON_ENGINEER
	case "2":
		mode = dboperation.ENGINEER
	default:
		log.Printf("invalid mode: %s", m)
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, `{"error":"invalid mode"}`)
		return
	}

	words, err := dboperation.SelectWords(mode, gameID)
	if err != nil {
		log.Printf("cannot select database: word")
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, `{"error":"`+err.Error()+`"}`)
		return
	}

	res := []models.WordGetResponse{}
	for _, word := range words {
		res = append(res, models.WordGetResponse{
			Word:    word.Word,
			Meaning: word.Meaning,
		})
	}

	fmt.Fprint(w, res)
}
