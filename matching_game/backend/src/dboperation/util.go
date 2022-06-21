package dboperation

import (
	"context"
	"log"
	"os"

	"github.com/SuperTikuwa/matching_game/models"
	"github.com/SuperTikuwa/matching_game/sheetclient"
	"github.com/joho/godotenv"
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

const DSN = DB_USER + ":" + DB_PASSWORD + "@tcp(" + DB_HOST + ":" + DB_PORT + ")/" + DB_NAME + "?parseTime=true"

func GormConnect() *gorm.DB {
	db, err := gorm.Open(mysql.Open(DSN), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	return db
}

func loadEnv() {
	if err := godotenv.Load("./.env"); err != nil {
		log.Fatal("Error loading .env file")
	}
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

	if err := db.AutoMigrate(&models.Difficulty{}, &models.Word{}, &models.Hash{}); err != nil {
		panic(err)
	}
}

// seed difficulty table
func init() {
	db := GormConnect()
	defer closeDB(db)

	var count int64 = 0
	db.Model(&models.Difficulty{}).Count(&count)

	if count > 0 {
		return
	}

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

// seed hashes table
func init() {
	db := GormConnect()
	defer closeDB(db)

	var count int64 = 0
	db.Model(&models.Hash{}).Count(&count)

	if count > 0 {
		return
	}

	difficulties := []models.Difficulty{}
	if err := db.Model(&models.Difficulty{}).Find(&difficulties).Error; err != nil {
		log.Fatal(err, "Cannot select database: difficulty")
	}

	hashes := []models.Hash{}
	for _, difficulty := range difficulties {
		hash := models.Hash{
			DifficultyID: difficulty.ID,
			Hash:         "",
		}
		hashes = append(hashes, hash)
	}

	if err := db.Create(&hashes).Error; err != nil {
		log.Fatal(err, "Cannot seed database: hash")
	}
}

// seed words table
func init() {
	loadEnv()

	db := GormConnect()
	defer closeDB(db)

	sc, err := sheetclient.NewSheetClient(context.Background(), os.Getenv("SPREAD_SHEET_ID"))
	if err != nil {
		log.Fatal(err, "Cannot instantiate sheet client")
	}

	hashes := []models.Hash{}
	if err := db.Model(&models.Hash{}).Find(&hashes).Error; err != nil {
		log.Fatal(err, "Cannot select database: hash")
	}

	for _, hash := range hashes {
		h, err := sc.GenerateHash(sheetclient.STUDENT)
		if err != nil {
			log.Fatal(err, "Cannot generate hash")
		}

		if hash.Hash != h {
			// if err := db.Model(&models.Word{}).Delete("difficulty_id = ?", hash.DifficultyID).Error; err != nil {
			// 	log.Fatal(err, "Cannot delete words")
			// }

			hash.Hash = h
			if err := db.Save(&hash).Error; err != nil {
				log.Fatal(err, "Cannot update hash")
			}

			sheet := sheetclient.IdToString(int(hash.ID))

			words, err := sc.Get(sheet)
			if err != nil {
				log.Fatal(err, "Cannot get words")
			}

			if err := db.Create(&words).Error; err != nil {
				log.Fatal(err, "Cannot seed database: word")
			}
		}
	}
}
