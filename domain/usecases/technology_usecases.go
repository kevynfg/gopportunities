package usecases

import (
	"fmt"

	"github.com/kevynfg/gopportunities/domain/models"
	"github.com/kevynfg/gopportunities/infra/repositories"
)

type TechnologiesUsecases struct {
	repository repositories.TechnologiesRepositorySql
}

type TechnologyInput struct {
	Name string	`json:"name"`
}

type TechnologyOutput struct {
	ID   uint	`json:"id"`
	Name string	`json:"name"`
}

func NewTechnologyUsecases(repository repositories.TechnologiesRepositorySql) *TechnologiesUsecases {
	return &TechnologiesUsecases{repository: repository}
}

func (u *TechnologiesUsecases) Execute(input models.TechnologyRequest) (*TechnologyOutput, error) {
	technology := models.NewTechnology(input.Name)
	result, err := u.repository.CreateTechnology(technology)
	if err != nil {
		return nil, err
	}

	return &TechnologyOutput{
		ID:   result.ID,
		Name: result.Name,
	}, nil
}

func (u *TechnologiesUsecases) FindAll() ([]*TechnologyOutput, error) {
	technologies, err := u.repository.FindAll()
	if err != nil {
		return nil, err
	}
	fmt.Println(technologies)
	var technologiesOutput []*TechnologyOutput
	for _, technology := range technologies {
		technologiesOutput = append(technologiesOutput, &TechnologyOutput{
			ID:   technology.ID,
			Name: technology.Name,
		})
	}

	return technologiesOutput, nil
}