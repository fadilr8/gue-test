package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

var SECRET string

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	SECRET = os.Getenv("JWT_SECRET")
	if SECRET == "" {
		SECRET = "default_value"
	}
}
