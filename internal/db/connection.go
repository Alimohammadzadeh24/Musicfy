package db

import (
	"log"
	"os"
	"sync"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	database *gorm.DB
	dbOnce   sync.Once
)

func InitializeDatabase() {
	dbOnce.Do(func() {
		err := godotenv.Load()
		if err != nil {
			log.Fatalf("Error loading .env file: %v", err)
		}

		dbUrl := os.Getenv("DATABASE_URL")
		db, err := gorm.Open(postgres.Open(dbUrl), &gorm.Config{})
		if err != nil {
			log.Fatalf("Failed to connect to database: %v", err)
		}
		database = db
	})
}

func GetDatabase() *gorm.DB {
	if database == nil {
		InitializeDatabase()
	}
	return database
}
