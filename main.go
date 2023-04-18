package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/julianNot/helpet-api-rest/routes"
	"github.com/julianNot/helpet-api-rest/models"
	"github.com/julianNot/helpet-api-rest/db"
)

func main() {
	err := godotenv.Load()
	if err != nil {
    log.Fatal("Error loading .env file")
  }
	host := os.Getenv("DB_HOST")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")
	// port := os.Getenv("PORT")
	port := os.Getenv("DB_PORT")
	db.DBConnection(host, user, password, dbname, port)

	db.DB.AutoMigrate(models.Partner{})
	db.DB.AutoMigrate(models.Vet{})
	db.DB.AutoMigrate(models.Professional{})
	db.DB.AutoMigrate(models.Attendant{})
	db.DB.AutoMigrate(models.Pet{})
	db.DB.AutoMigrate(models.Post{})
	db.DB.AutoMigrate(models.Quote{})

	router := mux.NewRouter()

	router.HandleFunc("/partners", routes.GetPartnersHandler).Methods("GET")
	router.HandleFunc("/partners/{id}", routes.GetPartnerHandler).Methods("GET")
	router.HandleFunc("/partners", routes.GetPartnerHandler).Methods("POST")
	router.HandleFunc("/partners/{id}", routes.EditPartnerHandler).Methods("PUT")
	router.HandleFunc("/partners/{id}", routes.DeletePartnerHandler).Methods("DELETE")

	http.ListenAndServe(":3000", router)
}
