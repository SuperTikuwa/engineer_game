package handler

import (
	"fmt"
	"log"
	"net/http"

	"github.com/SuperTikuwa/matching_game/dboperation"
	"github.com/SuperTikuwa/matching_game/models"
)

func WordsGetHandler(w http.ResponseWriter, r *http.Request) {
	words, err := dboperation.SelectWords(dboperation.STUDENT)
	if err != nil {
		log.Fatal(err)
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
