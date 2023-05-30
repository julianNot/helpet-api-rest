package models

import "gorm.io/gorm"

type Appointment struct {
	gorm.Model
	Service        string `json:"service_appointment"`
	Date           string `json:"date_appointment"`
	Name           string `json:"name_pet_appointment"`
	Male           string `json:"male_pet_appointment"`
	Description    string `json:"description_pet_appointment"`
	NameAttendant  string `json:"name_attendant_appointment"`
	Telephone      string `json:"telephone_attendant_appointment"`
	VetID          uint   `json:"vet_id_appointment"`
	ProfessionalID uint   `json:"professional_id_appointment"`
	AttendantID    uint		`json:"attendant_id_appointment"`
	NameAtte       string `json:"name_attendant"`
}
