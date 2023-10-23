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

func (h *OpportunitiesUsecases) CreateOpportunity(input models.OpportunityRequest) (*models.OpportunityResponse, error) {
	opportunity := models.NewOpportunity(input)
	result, err := h.repository.CreateOpportunity(opportunity)
	if err != nil {
		return &models.OpportunityResponse{}, err
	}

	return result, nil
}

func (h *OpportunitiesUsecases) EditOpportunity(input models.Opportunity) (*models.OpportunityResponse, error) {
	opportunity := models.EditOpportunity(input)
	result, err := h.repository.EditOpportunity(*opportunity)
	if err != nil {
		return &models.OpportunityResponse{}, err
	}

	return result, nil
}

func (h *OpportunitiesUsecases) FindAll(limit string, offset string) ([]models.OpportunityResponse, error) {
	opportunities, err := h.repository.FindAll(limit, offset)
	if err != nil {
		return nil, err
	} 
	var oppotunitiesOutput []models.OpportunityResponse
	for _, opportunity := range opportunities {
		oppotunitiesOutput = append(oppotunitiesOutput, models.OpportunityResponse{
			ID: opportunity.ID,
			CompanyName: opportunity.CompanyName,
			TechnologyNames: opportunity.TechnologyNames,
			Remote: opportunity.Remote,
			ContractType: opportunity.ContractType,
			Description: opportunity.Description,
			Link: opportunity.Link,
			Location: opportunity.Location,
			Name: opportunity.Name,
			Salary: opportunity.Salary,
			CreatedAt: opportunity.CreatedAt,
			Active: opportunity.Active,
		})
	}
	return oppotunitiesOutput, nil
}