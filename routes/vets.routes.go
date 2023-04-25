package routes

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/julianNot/helpet-api-rest/db"
	"github.com/julianNot/helpet-api-rest/models"
)

func GetVetsHandler(w http.ResponseWriter, r *http.Request) {
	var vets []models.Vet
	db.DB.Find(&vets)
	json.NewEncoder(w).Encode(vets)
}

func GetVetHandler(w http.ResponseWriter, r *http.Request) {
	var vet models.Professional
	params := mux.Vars(r)
	db.DB.First(&vet, params["id"])
	if vet.ID == 0 {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Vet Not Found"))
		return
	}
	json.NewEncoder(w).Encode(&vet)
}

func CreateVetHandler(w http.ResponseWriter, r *http.Request) {
	var vet models.Vet
	json.NewDecoder(r.Body).Decode(&vet)
	createdVet := db.DB.Create(&vet)
	err := createdVet.Error
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}
	json.NewEncoder(w).Encode(&vet)
}