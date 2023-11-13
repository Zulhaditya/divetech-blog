package initdb

import (
	"github.com/joho/godotenv"
	"log"
)

// LoadEnv get environment variable from .env file
func LoadEnv() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Unable to load .env file")
	}
}
