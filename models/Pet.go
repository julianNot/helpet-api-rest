package models

import "gorm.io/gorm"

type Pet struct {
	gorm.Model
	Name        string
	Male        bool
	Description string
	Image       string
	AttendantID uint
	PetID       uint
}
