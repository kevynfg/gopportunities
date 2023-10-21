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
	Name string `gorm:"unique"`
	Stack uint
	CreatedAt time.Time
	UpdatedAt time.Time
}

type TechnologyResponse struct {
	ID  uint `json:"id"`
	Name string `json:"name"`
	Stack uint `json:"stack"`
}

type TechnologyRequest struct {
	Name string `json:"name"`
	Stack uint `json:"stack"`
}

type OpportunityTechnology struct {
	TechnologyId uint
}

func NewTechnology(name string, stack uint) *Technology {
	return &Technology{
		Name: name,
		Stack: stack,
	}
}