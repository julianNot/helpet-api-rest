package services

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/julianNot/helpet-api-rest/db"
	"github.com/julianNot/helpet-api-rest/models"
)

func GetAppointmentsHandler(w http.ResponseWriter, r *http.Request) {
	var appointment []models.Appointment
	db.DB.Find(&appointment)
	json.NewEncoder(w).Encode(appointment)
}

func GetAppointmentHandler(w http.ResponseWriter, r *http.Request) {
	var appointment models.Appointment
	params := mux.Vars(r)
	db.DB.First(&appointment, params["id"])
	if appointment.ID == 0 {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Appointment Not Found"))
		return
	}
	json.NewEncoder(w).Encode(&appointment)
}

func GetAttendantsByAttendant(w http.ResponseWriter, r *http.Request) {
	var attendant models.Attendant
	params := mux.Vars(r)
	db.DB.Preload("Appointment").First(&attendant, params["id"])
	if attendant.ID == 0 {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Appointment Not Found"))
		return
	}
	db.DB.Find(&attendant.Appointment)
	if attendant.Appointment == nil {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Appointment Not Found"))
		return
	}
	json.NewEncoder(w).Encode(&attendant.Appointment)
}


func CreateAppointmentHandler(w http.ResponseWriter, r *http.Request) {
	var appointment models.Appointment
	json.NewDecoder(r.Body).Decode(&appointment)
	createdAppointment := db.DB.Create(&appointment)
	err := createdAppointment.Error
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}
	json.NewEncoder(w).Encode(&appointment)
}