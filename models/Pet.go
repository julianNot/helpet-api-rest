package models

import "gorm.io/gorm"

type Pet struct {
	gorm.Model
	Name          string `gorm:"not null" json:"name_pet"`
	Male          bool   `gorm:"not null" json:"male_pet"`
	Description   string `json:"description_pet"`
	Image         string `gorm:"not null" json:"img_pet"`
	// AttendantID   uint   `gorm:"not null" json:"attendant_id_pet"`
	AppointmentID uint   `gorm:"not null" json:"appointment_id_pet"`
}
