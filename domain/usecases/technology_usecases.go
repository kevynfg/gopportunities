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
	Stack uint `json:"stack"`
}

type TechnologyOutput struct {
	ID   uint	`json:"id"`
	Name string	`json:"name"`
	Stack uint `json:"stack"`
}

func NewTechnologyUsecases(repository repositories.TechnologiesRepositorySql) *TechnologiesUsecases {
	return &TechnologiesUsecases{repository: repository}
}

func (u *TechnologiesUsecases) Execute(input models.TechnologyRequest) (*TechnologyOutput, error) {
	technology := models.NewTechnology(input.Name, input.Stack)
	result, err := u.repository.CreateTechnology(technology)
	if err != nil {
		return nil, err
	}

	return &TechnologyOutput{
		ID:   result.ID,
		Name: result.Name,
		Stack: result.Stack,
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
			Stack: technology.Stack,
		})
	}

	return technologiesOutput, nil
}