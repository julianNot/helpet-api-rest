package helpers

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func GetPortServer() string {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	return os.Getenv("PORT")
}
