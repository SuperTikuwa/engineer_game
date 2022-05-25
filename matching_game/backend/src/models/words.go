package models

import "gorm.io/gorm"

type Words struct {
	gorm.Model
	ID         int        `gorm:"primary_key"`
	Word       string     `gorm:"not null"`
	Meaning    string     `gorm:"not null"`
	Difficulty Difficulty `gorm:"not null"`
}

type Difficulty struct {
	gorm.Model
	ID         int    `gorm:"primary_key"`
	Difficulty string `gorm:"not null"`
}
