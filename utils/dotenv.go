package utils

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

// GetDotEnvValue read .env file and return the value of the key
func GetDotEnvValue(key string) string {
	// Load .env file
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading .env file")
	}

	return os.Getenv(key)
}
