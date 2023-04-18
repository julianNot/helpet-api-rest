package models

import "gorm.io/gorm"

type Post struct {
	gorm.Model
	Title       string
	Subtitle    string
	Description string
	Imagen      string
	VetID       uint
	AttendantID   uint
}
