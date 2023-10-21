package factories

import (
	"github.com/kevynfg/gopportunities/application/controllers"
	"github.com/kevynfg/gopportunities/domain/usecases"
	"github.com/kevynfg/gopportunities/infra/repositories"
)

func NewTechnologyController(technologyRepository repositories.TechnologiesRepositorySql) *controllers.TechnologyHandler {
	technologyUsecase := usecases.NewTechnologyUsecases(technologyRepository)
	return controllers.NewTechnologyHandler(*technologyUsecase)
}