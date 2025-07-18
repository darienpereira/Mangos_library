package database

import (
	"fmt"
	"library/models"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var Db *gorm.DB

func OpenDb() {
	// err := godotenv.Load()
	// if err != nil {
	// 	log.Fatal("failed to load .env file", err)
	// }
	// connStr := "user=postgres password=password dbname=library port=5432 sslmode=disable"
	connStr := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PORT"),
	)

	Db, err = gorm.Open(postgres.Open(connStr), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to open database", err)
	}

	err = Db.Exec(`CREATE EXTENSION IF NOT EXISTS "uuid-ossp"`).Error
	if err != nil {
		log.Fatal("Failed to enable uuid-ossp extension:", err)
	}

	err = Db.AutoMigrate(&models.User{}, &models.Book{})
	if err != nil {
		log.Fatal("failed to migrate schema", err)
	}

	fmt.Println("connected to database successfully")

}