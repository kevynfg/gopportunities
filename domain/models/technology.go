package models

import (
	"time"

	"gorm.io/gorm"
)

type TechnologyRepository interface {
	CreateTechnology(technology *Technology) (Technology, error)
	FindAll() ([]*Technology, error)
}

type Technology struct {
	gorm.Model
	ID  uint `gorm:"primaryKey"`
	Name string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type TechnologyResponse struct {
	ID  uint `json:"id"`
	Name string `json:"name"`
}

type TechnologyRequest struct {
	Name string `json:"name"`
}

func NewTechnology(name string) *Technology {
	return &Technology{
		Name: name,
	}
}