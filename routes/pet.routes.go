package routes

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/julianNot/helpet-api-rest/db"
	"github.com/julianNot/helpet-api-rest/models"
)

func GetPetsHandler(w http.ResponseWriter, r *http.Request) {
	var pets []models.Pet
	db.DB.Find(&pets)
	json.NewEncoder(w).Encode(pets)
}

func GetPetHandler(w http.ResponseWriter, r *http.Request) {
	var pet models.Pet
	params := mux.Vars(r)
	db.DB.First(&pet, params["id"])
	if pet.ID == 0 {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Pet Not Found"))
		return
	}
	json.NewEncoder(w).Encode(&pet)
}

func CreatePetHandler(w http.ResponseWriter, r *http.Request) {
	var pet models.Pet
	json.NewDecoder(r.Body).Decode(&pet)
	createdPet:= db.DB.Create(&pet)
	err := createdPet.Error
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}
	json.NewEncoder(w).Encode(&pet)
}