package dboperation

import (
	"fmt"

	"github.com/SuperTikuwa/matching_game/models"
)

const (
	STUDENT      = "student"
	NON_ENGINEER = "non_engineer"
	ENGINEER     = "engineer"
)

func SelectWords(mode string) {
	db := GormConnect()
	defer closeDB(db)

	words := []models.Word{}
	db.Model(&models.Word{}).Order("NEWID()").Limit(10).Find(&words)

	fmt.Println(words)

}
