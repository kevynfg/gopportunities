package factories

import (
	"github.com/kevynfg/gopportunities/application/controllers"
	"github.com/kevynfg/gopportunities/domain/usecases"
	"github.com/kevynfg/gopportunities/infra/repositories"
)

func NewCompanyController(companyRepository repositories.CompaniesRepositorySql) *controllers.CompanyHandler {
	companyUsecase := usecases.NewCompanyUsecases(companyRepository)
	return controllers.NewCompanyHandler(*companyUsecase)
}