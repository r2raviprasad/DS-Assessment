package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

var (
	Username  string
	Password  string
	JwtSecret string
)

func LoadConfig() {
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found. Loading config from environment variable.")
	}

	Username = getEnv("APP_USERNAME", "strak")
	Password = getEnv("APP_PASSWORD", "strak")
	JwtSecret = getEnv("JWT_SECRET", "strak")
}

func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}
