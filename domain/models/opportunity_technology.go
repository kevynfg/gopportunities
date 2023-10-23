package models

import (
	"gorm.io/gorm"
)

type Opportunity_Technology struct {
	gorm.Model
	OpportunityID uint	`gorm:"foreignkey:OpportunityID"`
	TechnologyID uint	`gorm:"foreignkey:TechnologyID"`
}

func NewOpportunityTechnology(opportunityId uint, technology uint) *Opportunity_Technology {
	return &Opportunity_Technology{
		OpportunityID: opportunityId,
		TechnologyID: technology,
	}
}