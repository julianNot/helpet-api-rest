package models

import (
	"gorm.io/gorm"
)

type Partner struct {
	gorm.Model
	Name        string `json:"name_partner"`
	NameCompany string `gorm:"not null" json:"name_company_partner"`
	Email       string `gorm:"not null" json:"email_partner"`
	Nit         uint   `gorm:"not null" json:"nit_partner"`
	IsActive    bool   `json:"isActive_partner"`
	Password    string `gorm:"not null" json:"password"`
	Username    string `json:"username"`
	Vets        []Vet  `json:"vets"`
}
