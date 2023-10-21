package models

import (
	"time"

	"gorm.io/gorm"
)

type Opportunity struct {
	gorm.Model
	ID 						uint 					`gorm:"primaryKey"`
	Name 					string
	Description 	string
	Company 			*Company	 		`gorm:"foreignKey:CompanyID"`
	CompanyID 		uint
	Location 			string
	Remote 				bool
	ContractType 	string
	Technologies 	[]*Opportunity_Technology	`gorm:"many2many:opportunity_technologies;"`
	Link 					string
	Salary 				float64
	CreatedAt 		time.Time
	UpdatedAt 		time.Time
}

type OpportunityResponse struct {
	ID 						uint 						`json:"id"`
	Name 					string 					`json:"name"`
	Description 	string 					`json:"description"`
	Company 			*CompanyResponse`json:"company"`
	Location 			string 					`json:"location"`
	Remote 				bool 						`json:"remote"`
	ContractType 	string 					`json:"contract_type"`
	Technologies 	[]map[string]string	`json:"technologies"`
	Link 					string 					`json:"link"`
	Salary 				float64 				`json:"salary"`
	CreatedAt 		string 					`json:"created_at"`
}

type OpportunityRequest struct {
	Name 					string 					`json:"name"`
	Description 	string 					`json:"description"`
	CompanyID 		uint 						`json:"company_id"`
	Location 			string 					`json:"location"`
	Remote 				bool 						`json:"remote"`
	ContractType 	string 					`json:"contract_type"`
	Technologies 	[]uint 	`json:"technologies"`
	Link 					string 					`json:"link"`
	Salary 				float64 				`json:"salary"`
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
	}
	return opportunity
}