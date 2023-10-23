package models

import (
	"time"

	"gorm.io/gorm"
)

type Opportunity struct {
	gorm.Model
	ID 						uint 					`gorm:"primaryKey"`
	Name 					string				`json:"name"`
	Description 	string				`json:"description"`
	Company 			Company	 		`gorm:"foreignKey:CompanyID"`
	CompanyID 		uint					`json:"company_id"`
	Location 			string				`json:"location"`
	Remote 				bool					`json:"remote"`
	ContractType 	string				`json:"contract_type"`
	Technologies 	[]*Opportunity_Technology	`gorm:"many2many:opportunity_technologies;"`
	Link 					string				`json:"link"`
	Salary 				float64			`json:"salary"`
	CreatedAt 		time.Time		`json:"created_at"`
	UpdatedAt 		time.Time		`json:"updated_at"`
	Active 				bool					`json:"active"`
}

type OpportunityResponse struct {
	ID 						uint 							`json:"id"`
	Name 					string 					`json:"name"`
	Description 	string 					`json:"description"`
	CompanyName			string					`json:"company_name"`
	Location 			string 					`json:"location"`
	Remote 				bool 						`json:"remote"`
	ContractType 	string 					`json:"contract_type"`
	TechnologyNames 	*string				`json:"technology_names"`
	Link 					string 					`json:"link"`
	Salary 				float64 					`json:"salary"`
	CreatedAt 		time.Time 				`json:"created_at"`
	Active 				bool 						`json:"active"`
}

type OpportunityRequest struct {
	ID 						uint 							`json:"id"`
	Name 					string 					`json:"name"`
	Description 	string 					`json:"description"`
	CompanyID 		uint 							`json:"company_id"`
	Location 			string 					`json:"location"`
	Remote 				bool 						`json:"remote"`
	ContractType 	string 					`json:"contract_type"`
	Technologies 	[]uint 						`json:"technologies"`
	Link 					string 					`json:"link"`
	Salary 				float64 					`json:"salary"`
	CreatedAt 		time.Time					`json:"created_at"`
	Active 				bool 						`json:"active"`
}

func NewOpportunity(input OpportunityRequest) *OpportunityRequest {
	opportunity := &OpportunityRequest{
		Name: input.Name,
		Description: input.Description,
		CompanyID: input.CompanyID,
		Location: input.Location,
		Remote: input.Remote,
		ContractType: input.ContractType,
		Link: input.Link,
		Salary: input.Salary,
		Technologies: input.Technologies,
		CreatedAt: time.Now(),
		Active:	true,
	}
	return opportunity
}

func EditOpportunity(input Opportunity) *Opportunity {
	opportunity := Opportunity{
		ID: input.ID,
		Name: input.Name,
		Description: input.Description,
		CompanyID: input.CompanyID,
		Location: input.Location,
		Remote: input.Remote,
		ContractType: input.ContractType,
		Link: input.Link,
		Salary: input.Salary,
		Technologies: input.Technologies,
		Active: input.Active,
	}
	return &opportunity
}