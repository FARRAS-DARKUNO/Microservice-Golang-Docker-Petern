package config

import (
	"fmt"
	"log"
	"os"

	"github.com/FARRAS-DARKUNO/library-management/book-service/models"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	dsn := os.Getenv("BOOK_DB_URL")

	if dsn == "" {
		log.Fatal("BOOK_DB_URL environment variable is not set")
	} else {
		fmt.Println(dsn)
	}

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	DB = db

	if err := db.AutoMigrate(&models.BookManagement{}); err != nil {
		log.Fatalf("Failed to migrate BookManagement model: %v", err)
	}
	if err := db.AutoMigrate(&models.BookStock{}); err != nil {
		log.Fatalf("Failed to migrate BookStock model: %v", err)
	}
	if err := db.AutoMigrate(&models.BorrowBook{}); err != nil {
		log.Fatalf("Failed to migrate BorrowBook model: %v", err)
	}

	fmt.Println("Database connected and User table migrated successfully!")
}
