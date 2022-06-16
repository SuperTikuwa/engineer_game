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

func SelectWords(mode string) ([]models.Word, error) {
	db := GormConnect()
	defer closeDB(db)

	var difficultyID int
	switch mode {
	case STUDENT:
		difficultyID = 1
	case NON_ENGINEER:
		difficultyID = 2
	case ENGINEER:
		difficultyID = 3
	default:
		return nil, fmt.Errorf("invalid mode: %s", mode)
	}

	words := []models.Word{}
	if db.Model(&models.Word{}).Where("difficulty_id=?", difficultyID).Order("RAND()").Limit(10).Find(&words).Error != nil {
		return nil, fmt.Errorf("Cannot select database: word")
	}

	return words, nil
}
