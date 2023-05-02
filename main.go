package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/julianNot/helpet-api-rest/routes"
	"github.com/julianNot/helpet-api-rest/auth"
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
	db.DB.AutoMigrate(models.Appointment{})

	router := mux.NewRouter()

	// Login Partners
	router.HandleFunc("/auth/partners", auth.AuthenticatePartnerHandler).Methods("POST")
	
	// Routes Partners
	router.HandleFunc("/partners", routes.GetPartnersHandler).Methods("GET")
	router.HandleFunc("/partners/{id}", routes.GetPartnerHandler).Methods("GET")
	router.HandleFunc("/partners", routes.CreatePartnerHandler).Methods("POST")
	router.HandleFunc("/partners/{id}", routes.UpdatePartnerHandler).Methods("PUT")
	router.HandleFunc("/partners/{id}", routes.DeletePartnerHandler).Methods("DELETE")

	// Routes Vets
	router.HandleFunc("/vets", routes.GetVetsHandler).Methods("GET")
	router.HandleFunc("/vets/{id}", routes.GetVetHandler).Methods("GET")
	router.HandleFunc("/vets", routes.CreateVetHandler).Methods("POST")

	// Routes Professionals
	router.HandleFunc("/professionals", routes.GetProfessionalsHandler).Methods("GET")
	router.HandleFunc("/professionals/{id}", routes.GetProfessionalHandler).Methods("GET")
	router.HandleFunc("/professionals", routes.CreateProfessionalHandler).Methods("POST")
	// TODO:

	// Routes Posts
	router.HandleFunc("/posts", routes.GetAllPostsHandler).Methods("GET")
	router.HandleFunc("/posts/{id}", routes.GetPostHandler).Methods("GET")
	router.HandleFunc("/posts", routes.CreatePostHandler).Methods("POST")
	// TODO:

	// Routes Pets
	router.HandleFunc("/pets", routes.GetPetsHandler).Methods("GET")
	router.HandleFunc("/pets/{id}", routes.GetPetHandler).Methods("GET")
	router.HandleFunc("/pets", routes.CreatePetHandler).Methods("POST")

	// Routes Attendants
	router.HandleFunc("/attendants", routes.GetAttendantsHandler).Methods("GET")
	router.HandleFunc("/attendants/{id}", routes.GetAttendantHandler).Methods("GET")
	router.HandleFunc("/attendants", routes.CreateAttendantHandler).Methods("POST")

	// Routes Appointments
	router.HandleFunc("/appointments", routes.GetAppointmentsHandler).Methods("GET")
	router.HandleFunc("/appointments/{id}", routes.GetAppointmentHandler).Methods("GET")
	router.HandleFunc("/appointments", routes.CreateAppointmentHandler).Methods("POST")

	http.ListenAndServe(":3000", router)
}
