package models

import (
	"gorm.io/gorm"
)

type Vet struct {
	gorm.Model
	NameVet          string         `gorm:"not null" json:"name_vet"`
	Location         string         `json:"location_vet"`
	City             string         `json:"city_vet"`
	Country          string         `json:"country_vet"`
	Presentation     string         `json:"presentation_vet"`
	Description      string         `json:"description_vet"`
	ShortDescription string         `json:"short_description_vet"`
	Image            string         `json:"image_vet"`
	Telephone        uint           `json:"telephone_vet"`
	Blog             []Post         `json:"blog_vet"`
	Professionals    []Professional `json:"profesionals_vet"`
	Appointment      []Appointment  `json:"appointment_vet"`
	PartnerID        uint           `json:"partner_id_vet"`
	BasicServices    string         `json:"basic_services_vet"`
	OptionalServices string         `json:"optional_services_vet"`
}
