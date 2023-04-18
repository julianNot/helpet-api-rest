package models

import "gorm.io/gorm"

type Quote struct {
	gorm.Model
	Service        string
	Date           string
	VetID          uint
	ProfessionalID uint
}
