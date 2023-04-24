package models

import "gorm.io/gorm"

type Attendant struct {
	gorm.Model
	Name        string
	Telephone   uint
	SaveBlogs   []Post
	Pet         []Pet
	Appointment []Appointment
}
