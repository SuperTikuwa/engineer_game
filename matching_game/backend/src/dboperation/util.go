package dboperation

import (
	"log"

	"github.com/SuperTikuwa/matching_game/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

const (
	DB_HOST     = "mariadb"
	DB_PORT     = "3306"
	DB_USER     = "golang"
	DB_PASSWORD = "my-secret-pw"
	DB_NAME     = "matching.db"
)

const DSN = DB_USER + ":" + DB_PASSWORD + "@tcp(" + DB_HOST + ":" + DB_PORT + ")/" + DB_NAME

func GormConnect() *gorm.DB {
	db, err := gorm.Open(mysql.Open(DSN), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	return db
}

func closeDB(db *gorm.DB) {
	d, err := db.DB()
	if err != nil {
		panic(err)
	}
	d.Close()
}

// migrate database schema
func init() {
	db := GormConnect()
	defer closeDB(db)

	if err := db.AutoMigrate(&models.Difficulty{}, &models.Words{}); err != nil {
		panic(err)
	}
}

// seed database
func init() {
	db := GormConnect()
	defer closeDB(db)

	difficulties := []models.Difficulty{
		{
			Difficulty: STUDENT,
		},
		{
			Difficulty: NON_ENGINEER,
		},
		{
			Difficulty: ENGINEER,
		},
	}

	if err := db.Create(&difficulties).Error; err != nil {
		log.Fatal(err, "Cannot seed database: difficulty")
	}

}
