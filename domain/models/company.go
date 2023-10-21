package models

import (
	"time"

	"gorm.io/gorm"
)

type Company struct {
	gorm.Model
	ID 	 uint `gorm:"primaryKey"`
	Name string
	Startup bool
	CreatedAt time.Time
	UpdatedAt time.Time
}

type CompanyResponse struct {
	ID 	 uint `json:"id"`
	Name string `json:"name"`
	Startup bool `json:"startup"`
	CreatedAt string `json:"created_at"`
}

type CompanyRequest struct {
	Name string `json:"name"`
	Startup bool `json:"startup"`
}

func NewCompany(name string, startup bool) *Company {
	return &Company{
		Name: name,
		Startup: startup,
	}
}