package models

import "gorm.io/gorm"

type Professional struct {
	gorm.Model
	Name      string
	Email     string
	Specialty string
	Telephone uint
	Quotes    []Quote
	VetID     uint
}
