package db

import (
	"BelajarGolang4/models"
	"log"
	"os"

	_ "database/sql"

	_ "github.com/lib/pq"

	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
)

func InitDB() *gorm.DB {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("error load env")
	}

	conn := os.Getenv("POSTGRES_URL")
	db, err := gorm.Open("postgres", conn)
	if err != nil {
		log.Fatal(err)
	}

	migrate(db)
	return db
}

func migrate(db *gorm.DB) {
	db.AutoMigrate(&models.Books{})

	data := models.Books{}
	if db.Find(&data).RecordNotFound() {
		seederBook(db)
	}
}

func seederBook(db *gorm.DB) {
	data := []models.Books{
		{
			Title:       "5",
			Author:      "gk ada",
			Description: "gk ada juga pd mls",
			Stock:       10,
		},
		{
			Title:       "payung",
			Author:      "gk ada juga",
			Description: "mls gk usah lh",
			Stock:       10,
		},
	}

	for _, v := range data {
		db.Create(&v)
	}
}
