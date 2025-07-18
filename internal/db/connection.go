package db

import (
	"database/sql"
	"log"
	"musicfy/internal/config"
	"sync"

	_ "github.com/lib/pq" // PostgreSQL driver
)

var (
	database *sql.DB
	dbOnce   sync.Once
)

// InitializeDatabase creates a database connection
func InitializeDatabase() {
	dbOnce.Do(func() {
		// Ensure config is loaded
		if config.AppConfig.DBConfig.URL == "" {
			config.LoadConfig()
		}

		dbURL := config.AppConfig.DBConfig.URL
		if dbURL == "" {
			log.Fatalf("Database URL not configured")
		}

		db, err := sql.Open("postgres", dbURL)
		if err != nil {
			log.Fatalf("Failed to connect to database: %v", err)
		}

		// Test the connection
		if err = db.Ping(); err != nil {
			log.Fatalf("Failed to ping database: %v", err)
		}

		// Configure connection pool based on environment
		db.SetMaxOpenConns(config.AppConfig.DBConfig.MaxConns)
		db.SetMaxIdleConns(config.AppConfig.DBConfig.IdleConns)

		database = db
		log.Println("Database connection established")
	})
}

// GetDB returns the database connection
func GetDB() *sql.DB {
	if database == nil {
		InitializeDatabase()
	}
	return database
}

// Close closes the database connection
func Close() error {
	if database != nil {
		return database.Close()
	}
	return nil
}
