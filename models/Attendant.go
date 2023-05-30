package models

import "gorm.io/gorm"

type Attendant struct {
	gorm.Model
	// Name        string        `gorm:"not null" json:"name_attendant"`
	Telephone   uint          `gorm:"not null" json:"telephone_attendant"`
	// SaveBlogs   []Post        `json:"saved_blogs_attendant"`
	// Pet         []Pet         `json:"pets_attendant"`
	Appointment []Appointment `json:"appointments_attendant"`
	NamePet     string        `gorm:"default:'Sin nombre'" json:"name_pet_attendant"`
	Username    string        `json:"username"`
	Password    string        `json:"password"`
}
