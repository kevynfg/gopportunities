package repositories

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/kevynfg/gopportunities/domain/models"
	"github.com/kevynfg/gopportunities/infra/logger"
	"github.com/kevynfg/gopportunities/infra/repositories/queries"
	"gorm.io/gorm"
)

type OpportunitiesRepositorySql struct {
	db *gorm.DB
}

func NewOpportunitiesRepositorySql(db *gorm.DB) *OpportunitiesRepositorySql {
	return &OpportunitiesRepositorySql{db: db}
}

func NewLog(p string) *logger.Logger {
	return logger.NewLogger(p)
}

func (r *OpportunitiesRepositorySql) CreateOpportunity(opportunity *models.OpportunityRequest) (models.Opportunity, error) {
	r.db.Exec(`
		CREATE TABLE IF NOT EXISTS opportunities (
				id INTEGER PRIMARY KEY,
				name TEXT,
				description TEXT,
				company_id INTEGER,
				location TEXT,
				remote BOOL,
				contract_type TEXT,
				link TEXT,
				salary FLOAT,
				created_at TIMESTAMP,
				updated_at TIMESTAMP
		)
	`)
	var newOpportunity models.Opportunity
	if err := r.db.Raw(`
    INSERT INTO opportunities (
        name,
        description,
        company_id,
        location,
        remote,
        contract_type,
        link,
        salary
				created_at
    ) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?)
    RETURNING id, name, description, company_id, location, remote, contract_type, link, salary, created_at
	`, opportunity.Name, opportunity.Description, opportunity.CompanyID, opportunity.Location, opportunity.Remote, opportunity.ContractType, opportunity.Link, opportunity.Salary, time.Now()).Scan(&newOpportunity).Error;
	err != nil {
			logger.NewLogger("repositories/opportunity_repository.go").Errf("Error creating opportunity: %v", err)
			return models.Opportunity{}, err
	}

	var opportunityTechnologies []interface{}
	for _, technology := range opportunity.Technologies {
			opportunityTechnologies = append(opportunityTechnologies, models.NewOpportunityTechnology(newOpportunity.ID, technology))
	}

	formattedOpportunityTechnologies := make([]interface{}, len(opportunityTechnologies))
	for _, value := range opportunityTechnologies {
    opportunityTechnology := *value.(*models.Opportunity_Technology)
		formattedOpportunityTechnologies = append(formattedOpportunityTechnologies, opportunityTechnology.OpportunityID, opportunityTechnology.TechnologyID)
	}
	
	query := `
			INSERT INTO opportunity_technologies (
					opportunity_id,
					opportunity_technology_id
			) VALUES %s
	`
	var valueStrings []string
	for i := 0; i < len(opportunityTechnologies)/2; i++ {
			valueStrings = append(valueStrings, "(?, ?)", "(?, ?)")
	}

	query = fmt.Sprintf(query, strings.Join(valueStrings, ","))
	var nonNilValues []interface{}
	for _, value := range formattedOpportunityTechnologies {
		if value != nil {
			nonNilValues = append(nonNilValues, value)
		}
	}
	query = fmt.Sprintf("%s ON CONFLICT DO NOTHING", query)	
	
	if errBatch := r.db.Exec(query, nonNilValues...).Error; errBatch != nil {
		logger.NewLogger("repositories/opportunity_repository.go").Errf("Error creating opportunity_technology: %v", errBatch)
		return models.Opportunity{}, errBatch
	}
	return newOpportunity, nil
}

func (r *OpportunitiesRepositorySql) EditOpportunity(input models.Opportunity) (*models.Opportunity, error) {
	opportunity := models.Opportunity(input)
	result := r.db.Model(&models.Opportunity{}).Where("id = ?", opportunity.ID).Updates(&opportunity)
	if result.Error != nil {
		return nil, result.Error
	}
	return &opportunity, nil
}

func (r *OpportunitiesRepositorySql) FindAll(limit string, offset string) ([]models.OpportunityResponse, error) {
	var opportunities []models.OpportunityResponse
	queryLimit, limitErr := strconv.Atoi(limit)
	if limitErr != nil {
		logger.NewLogger("repositories/opportunity_repository.go").Errf("Error converting limit to int: %v", limitErr)
		return nil, limitErr
	}
	queryOffset, offsetErr :=strconv.Atoi(offset)
	if offsetErr != nil {
		logger.NewLogger("repositories/opportunity_repository.go").Errf("Error converting offset to int: %v", offsetErr)
		return nil, offsetErr
	}
	if err := r.db.Raw(queries.FindAllOpportunitiesQuery(queryLimit, queryOffset)).Scan(&opportunities).Error;err != nil {
		logger.NewLogger("repositories/opportunity_repository.go").Errf("Error finding opportunities: %v", err)
		return nil, err
	}
	return opportunities, nil
}