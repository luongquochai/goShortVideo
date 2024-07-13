package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/luongquochai/goShortVideo/internal/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	var err error
	dsn := os.Getenv("DATABASE_URL")
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}
	// Run database migrations
	err = DB.AutoMigrate(&models.User{}, &models.Video{})
	if err != nil {
		log.Fatal("Failed to run database migrations:", err)
	}
}

func LoadEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading environment:", err)
	}
}
