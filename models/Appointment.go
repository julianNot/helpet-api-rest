package models

import "gorm.io/gorm"

type Appointment struct {
	gorm.Model
	Service        string
	Date           string
	VetID          uint
	ProfessionalID uint
	AttendantID    uint
}
