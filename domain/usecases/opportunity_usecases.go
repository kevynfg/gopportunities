package usecases

import (
	"github.com/kevynfg/gopportunities/domain/models"
	"github.com/kevynfg/gopportunities/infra/repositories"
)

type OpportunitiesUsecases struct {
	repository repositories.OpportunitiesRepositorySql
}

func NewOpportunitiesUsecases(repository repositories.OpportunitiesRepositorySql) *OpportunitiesUsecases {
	return &OpportunitiesUsecases{repository: repository}
}

type OpportunityInput struct {
	Name 					string 					`json:"name"`
	Description 	string 					`json:"description"`
	CompanyID 		uint 						`json:"company_id"`
	Location 			string 					`json:"location"`
	Remote 				bool 						`json:"remote"`
	ContractType 	string 					`json:"contract_type"`
	Technologies 	[]*models.Opportunity_Technology				`json:"technologies"`
	Link 					string 					`json:"link"`
	Salary 				float64 				`json:"salary"`
}

type OpportunityOutput struct {
	Name 					string 					`json:"name"`
	Description 	string 					`json:"description"`
	CompanyID 		uint 						`json:"company_id"`
	Location 			string 					`json:"location"`
	Remote 				bool 						`json:"remote"`
	ContractType 	string 					`json:"contract_type"`
	Technologies 	[]*models.Opportunity_Technology 	`json:"technologies"`
	Link 					string 					`json:"link"`
	Salary 				float64 				`json:"salary"`
}

func (h *OpportunitiesUsecases) CreateOpportunity(input models.OpportunityRequest) (models.Opportunity, error) {
	opportunity := models.NewOpportunity(input)
	result, err := h.repository.CreateOpportunity(opportunity)
	if err != nil {
		return models.Opportunity{}, err
	}

	return result, nil
}