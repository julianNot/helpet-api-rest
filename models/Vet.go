package models

import (
	"gorm.io/gorm"
)

type Vet struct {
	gorm.Model
	NameVet          string         `gorm:"not null" json:"name_vet"`
	Location         string         `gorm:"not null" json:"location_vet"`
	City             string         `gorm:"not null" json:"city_vet"`
	Country          string         `gorm:"not null" json:"country_vet"`
	Presentation     string         `gorm:"not null" json:"presentation_vet"`
	Description      string         `json:"description_vet"`
	ShortDescription string         `json:"short_description_vet"`
	Image            string         `json:"image_vet"`
	Telephone        uint           `gorm:"not null" json:"telephone_vet"`
	Blog             []Post         `json:"blog_vet"`
	Professionals    []Professional `json:"profesionals_vet"`
	Quotes           []Quote        `json:"quotes_vet"`
	PartnerID        uint						`json:"partner_id_vet"`
}
