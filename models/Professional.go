package models

import "gorm.io/gorm"

type Professional struct {
	gorm.Model
	Name        string        `gorm:"not null" json:"name_professional"`
	Email       string        `json:"email_professional"`
	Specialty   string        `gorm:"not null" json:"specialty_professional"`
	Telephone   uint          `gorm:"not null" json:"telephone_professional"`
	Appointment []Appointment `json:"appointment_professional"`
	VetID       uint          `gorm:"not null" json:"vet_id_professional"`
}
