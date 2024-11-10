package config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"github.com/FARRAS-DARKUNO/library-management/user-service/models"
)

var DB *gorm.DB

func InitDB() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	dsn := os.Getenv("USER_DB_URL")  

	if dsn == "" {
		log.Fatal("USER_DB_URL environment variable is not set")
	}else {
        fmt.Println(dsn)
    }

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	DB = db

	if err := db.AutoMigrate(&models.User{}); err != nil {
		log.Fatalf("Failed to migrate User model: %v", err)
	}

	fmt.Println("Database connected and User table migrated successfully!")
}