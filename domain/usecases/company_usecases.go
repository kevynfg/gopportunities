package usecases

import (
	"github.com/kevynfg/gopportunities/domain/models"
	"github.com/kevynfg/gopportunities/infra/repositories"
)

type CompaniesUsecases struct {
	repository repositories.CompaniesRepositorySql
}

type CompanyOutput struct {
	ID   uint	`json:"id"`
	Name string	`json:"name"`
	Startup bool `json:"startup"`
	CreatedAt string `json:"created_at"`
}

func NewCompanyUsecases(repository repositories.CompaniesRepositorySql) *CompaniesUsecases {
	return &CompaniesUsecases{repository: repository}
}

func (u *CompaniesUsecases) Execute(input models.CompanyRequest) (*CompanyOutput, error) {
	company := models.NewCompany(&input.Name, &input.Startup)
	result, err := u.repository.CreateCompany(company)
	if err != nil {
		return nil, err
	}

	return &CompanyOutput{
		ID:   result.ID,
		Name: *result.Name,
		Startup: *result.Startup,
		CreatedAt: result.CreatedAt.Local().String(),
	}, nil
}

func (u *CompaniesUsecases) FindAll() ([]*CompanyOutput, error) {
	companies, err := u.repository.FindAll()
	if err != nil {
		return nil, err
	}
	var companiesOutput []*CompanyOutput
	for _, company := range companies {
		companiesOutput = append(companiesOutput, &CompanyOutput{
			ID:   company.ID,
			Name: *company.Name,
			Startup: *company.Startup,
			CreatedAt: company.CreatedAt.Local().UTC().Format("2006-01-02 15:04:05"),
		})
	}

	return companiesOutput, nil
}