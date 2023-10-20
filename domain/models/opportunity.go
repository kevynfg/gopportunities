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
	Technologies 	[]*Technology `gorm:"many2many:opportunity_technologies;"`
	Link 					string
	Salary 				float64
	CreatedAt 		time.Time
}

type OpportunityResponse struct {
	ID 						uint 						`json:"id"`
	Name 					string 					`json:"name"`
	Description 	string 					`json:"description"`
	Company 			*CompanyResponse`json:"company"`
	Location 			string 					`json:"location"`
	Remote 				bool 						`json:"remote"`
	ContractType 	string 					`json:"contract_type"`
	Technologies 	[]*Technology 	`json:"technologies"`
	Link 					string 					`json:"link"`
	Salary 				float64 				`json:"salary"`
	CreatedAt 		string 					`json:"created_at"`
}