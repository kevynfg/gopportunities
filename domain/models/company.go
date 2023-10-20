package models

import "gorm.io/gorm"

type Company struct {
	gorm.Model
	ID 	 uint `gorm:"primaryKey"`
	Name string
}

type CompanyResponse struct {
	ID 	 uint `json:"id"`
	Name string `json:"name"`
}