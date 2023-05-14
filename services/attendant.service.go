package services

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/julianNot/helpet-api-rest/db"
	"github.com/julianNot/helpet-api-rest/models"
)

func GetAttendantsHandler(w http.ResponseWriter, r *http.Request) {
	var attendant []models.Attendant
	db.DB.Preload("Pet").Find(&attendant)
	json.NewEncoder(w).Encode(attendant)
}

func GetAttendantHandler(w http.ResponseWriter, r *http.Request) {
	var attendant models.Attendant
	params := mux.Vars(r)
	db.DB.First(&attendant, params["id"])
	if attendant.ID == 0 {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Attendant Not Found"))
		return
	}
	json.NewEncoder(w).Encode(&attendant)
}

func CreateAttendantHandler(w http.ResponseWriter, r *http.Request) {
	var attendant models.Attendant
	json.NewDecoder(r.Body).Decode(&attendant)
	createdAttendant := db.DB.Create(&attendant)
	err := createdAttendant.Error
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}
	json.NewEncoder(w).Encode(&attendant)
}