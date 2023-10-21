package factories

import (
	"github.com/kevynfg/gopportunities/application/controllers"
	"github.com/kevynfg/gopportunities/domain/usecases"
	"github.com/kevynfg/gopportunities/infra/repositories"
)

func NewOpportunityController(opportunityRepository repositories.OpportunitiesRepositorySql) *controllers.OpportunityHandler {
	opportunityUsecase := usecases.NewOpportunitiesUsecases(opportunityRepository)
	return controllers.NewOpportunityHandler(*opportunityUsecase)
}