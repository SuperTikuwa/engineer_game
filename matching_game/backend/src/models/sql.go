package models

import "gorm.io/gorm"

type Word struct {
	gorm.Model
	Word         string `gorm:"not null"`
	Meaning      string `gorm:"not null"`
	DifficultyID uint   `gorm:"not null"`
}

type Difficulty struct {
	gorm.Model
	Difficulty string `gorm:"not null"`
}

type Hash struct {
	gorm.Model
	Hash         string `gorm:"not null"`
	DifficultyID uint   `gorm:"not null"`
}
