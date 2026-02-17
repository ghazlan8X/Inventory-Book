package db

import (
	"BelajarGolang5/models"
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
	db.AutoMigrate(&models.User{})

	BooksData := models.Books{}
	UserData := models.User{}

	if db.Find(&BooksData).RecordNotFound() {
		seederBook(db)
	}

	if db.Find(&UserData).RecordNotFound() {
		seederBook(db)
	}
}

func seederBook(db *gorm.DB) {
	Booksdata := []models.Books{
		{
			Title:       "5cm Per Second",
			Author:      "gk ada",
			Description: "gk ada juga pd mls",
			Stock:       10,
		},
		{
			Title:       "5",
			Author:      "gk ada juga",
			Description: "mls gk usah lh",
			Stock:       10,
		},
	}

	UserData := []models.User{
		{
			Username: "admin",
			Password: "admin123",
		},
		{
			Username: "zhongli",
			Password: "123",
		},
	}

	for _, v := range Booksdata {
		db.Create(&v)
	}

	for _, v := range UserData {
		db.Create(&v)
	}
}
