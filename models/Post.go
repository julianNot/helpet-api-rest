package models

import "gorm.io/gorm"

type Post struct {
	gorm.Model
	Title       string `gorm:"not null" json:"title_post"`
	Subtitle    string `json:"subtitle_post"`
	Description string `gorm:"not null" json:"description_post"`
	Imagen      string `json:"imagen_post"`
	VetID       uint   `json:"vet_id_post"`
	AttendantID uint   `json:"attendant_id_post"`
}
