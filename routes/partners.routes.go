package routes

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/julianNot/helpet-api-rest/db"
	"github.com/julianNot/helpet-api-rest/models"
	"golang.org/x/crypto/bcrypt"
)

func GetPartnersHandler(w http.ResponseWriter, r *http.Request) {
	var partners []models.Partner
	db.DB.Find(&partners)
	json.NewEncoder(w).Encode(partners)
}

func GetPartnerHandler(w http.ResponseWriter, r *http.Request) {
	var partner models.Partner
	params := mux.Vars(r)
	db.DB.First(&partner, params["id"])
	if partner.ID == 0 {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Partner Not Found"))
		return
	}
	json.NewEncoder(w).Encode(&partner)
}

func CreatePartnerHandler(w http.ResponseWriter, r *http.Request) {
	var partner models.Partner
	json.NewDecoder(r.Body).Decode(&partner)
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(partner.Password), bcrypt.DefaultCost)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Failed to hash password"))
		return
	}
	partner.Password = string(hashedPassword)
	createdPartner := db.DB.Create(&partner)
	err = createdPartner.Error
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}
	json.NewEncoder(w).Encode(&partner)
}

func UpdatePartnerHandler(w http.ResponseWriter, r *http.Request) {
	var partner models.Partner
	params := mux.Vars(r)
	db.DB.First(&partner, params["id"])
	if partner.ID == 0 {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Partner Not Found"))
		return
	}
	json.NewDecoder(r.Body).Decode(&partner)
	db.DB.Save(&partner)
	json.NewEncoder(w).Encode(&partner)
}

func DeletePartnerHandler(w http.ResponseWriter, r *http.Request) {
	var partner models.Partner
	params := mux.Vars(r)
	db.DB.First(&partner, params["id"])
	if partner.ID == 0 {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Task Not Found"))
		return
	}
	db.DB.Unscoped().Delete(&partner)
	w.WriteHeader(http.StatusNoContent)
}
