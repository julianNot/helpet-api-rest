package models

import "gorm.io/gorm"

type Appointment struct {
	gorm.Model
	Service        string `gorm:"not null" json:"service_appointment"`
	Date           string `gorm:"not null" json:"date_appointment"`
	VetID          uint   `json:"vet_id_appointment"`
	ProfessionalID uint   `gorm:"not null" json:"professional_id_appointment"`
	AttendantID    uint   `gorm:"not null" json:"attendant_id_appointment"`
}
