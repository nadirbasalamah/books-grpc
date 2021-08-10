package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

// Config untuk mendapat nilai dari env
func Config(key string) string {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file\n")
	}
	return os.Getenv(key)
}
