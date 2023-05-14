package db

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/julianNot/helpet-api-rest/models"
)

func SetupDatabase() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	host := os.Getenv("DB_HOST")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")
	port := os.Getenv("DB_PORT")

	DBConnection(host, user, password, dbname, port)

	DB.AutoMigrate(
		&models.Partner{},
		&models.Vet{},
		&models.Professional{},
		&models.Attendant{},
		&models.Pet{},
		&models.Post{},
		&models.Appointment{},
	)
}
