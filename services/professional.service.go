package services

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/julianNot/helpet-api-rest/db"
	"github.com/julianNot/helpet-api-rest/models"
)

func GetProfessionalsHandler(w http.ResponseWriter, r *http.Request) {
	var professionals []models.Professional
	db.DB.Find(&professionals)
	json.NewEncoder(w).Encode(professionals)
}

func GetProfessionalHandler(w http.ResponseWriter, r *http.Request) {
	var professional models.Professional
	params := mux.Vars(r)
	db.DB.First(&professional, params["id"])
	if professional.ID == 0 {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Professional Not Found"))
		return
	}
	json.NewEncoder(w).Encode(&professional)
}

func CreateProfessionalHandler(w http.ResponseWriter, r *http.Request) {
	var professional models.Professional
	json.NewDecoder(r.Body).Decode(&professional)
	createdProfessional := db.DB.Create(&professional)
	err := createdProfessional.Error
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}
	json.NewEncoder(w).Encode(&professional)
}