package config

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/joho/godotenv"
)

// Environment represents the application environment
type Environment string

const (
	// Development environment
	Development Environment = "development"
	// Production environment
	Production Environment = "production"
	// Testing environment
	Testing Environment = "testing"
)

// Config holds the application configuration
type Config struct {
	Environment  Environment
	DBConfig     DatabaseConfig
	ServerConfig ServerConfig
	JWTConfig    JWTConfig
}

// DatabaseConfig holds database configuration
type DatabaseConfig struct {
	URL       string
	MaxConns  int
	IdleConns int
}

// ServerConfig holds server configuration
type ServerConfig struct {
	Port string
	Host string
}

// JWTConfig holds JWT configuration
type JWTConfig struct {
	Secret      string
	ExpiryHours int
}

var (
	// AppConfig is the global application configuration
	AppConfig Config
)

// LoadConfig loads configuration from .env file and environment variables
func LoadConfig() {
	// Load .env file if it exists
	loadEnvFile()

	// Determine environment
	env := getEnvironment()

	// Set up configuration based on environment
	AppConfig = Config{
		Environment: env,
		DBConfig: DatabaseConfig{
			URL:       getEnv("DATABASE_URL", ""),
			MaxConns:  getEnvAsInt("DB_MAX_CONNS", 25),
			IdleConns: getEnvAsInt("DB_IDLE_CONNS", 5),
		},
		ServerConfig: ServerConfig{
			Port: getEnv("APP_PORT", "8080"),
			Host: getEnv("APP_HOST", "0.0.0.0"),
		},
		JWTConfig: JWTConfig{
			Secret:      getEnv("JWT_SECRET", "default_secret_change_in_production"),
			ExpiryHours: getEnvAsInt("JWT_EXPIRY_HOURS", 24),
		},
	}

	// Log the current environment
	log.Printf("Application running in %s mode", env)

	// Validate critical configuration
	validateConfig()
}

// getEnvironment determines the current environment
func getEnvironment() Environment {
	env := strings.ToLower(os.Getenv("APP_ENV"))

	switch env {
	case "production", "prod":
		return Production
	case "testing", "test":
		return Testing
	default:
		return Development
	}
}

// loadEnvFile loads the .env file if it exists
func loadEnvFile() {
	// Try to load .env file
	if err := godotenv.Load(); err != nil {
		log.Printf("No .env file found, using environment variables")
	} else {
		log.Printf("Loaded configuration from .env file")
	}
}

// validateConfig validates critical configuration settings
func validateConfig() {
	// In production, ensure we have a proper JWT secret
	if IsProduction() && (AppConfig.JWTConfig.Secret == "" ||
		AppConfig.JWTConfig.Secret == "default_secret_change_in_production") {
		log.Fatalf("Production environment requires a secure JWT_SECRET to be set")
	}

	// Ensure database URL is set
	if AppConfig.DBConfig.URL == "" {
		log.Printf("Warning: DATABASE_URL is not set")
	}
}

// IsDevelopment checks if the current environment is development
func IsDevelopment() bool {
	return AppConfig.Environment == Development
}

// IsProduction checks if the current environment is production
func IsProduction() bool {
	return AppConfig.Environment == Production
}

// IsTesting checks if the current environment is testing
func IsTesting() bool {
	return AppConfig.Environment == Testing
}

// Helper functions to get environment variables with defaults
func getEnv(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}

func getEnvAsInt(key string, defaultValue int) int {
	valueStr := os.Getenv(key)
	if value, err := fmt.Sscanf(valueStr, "%d", &defaultValue); err != nil || value == 0 {
		return defaultValue
	}
	return defaultValue
}
