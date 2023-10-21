package repositories

import (
	"fmt"

	"github.com/kevynfg/gopportunities/domain/models"
	"gorm.io/gorm"
)

type OpportunitiesRepositorySql struct {
	db *gorm.DB
}

func NewOpportunitiesRepositorySql(db *gorm.DB) *OpportunitiesRepositorySql {
	return &OpportunitiesRepositorySql{db: db}
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
	err := r.db.Raw(`
    INSERT INTO opportunities (
        name,
        description,
        company_id,
        location,
        remote,
        contract_type,
        link,
        salary
    ) VALUES (?, ?, ?, ?, ?, ?, ?, ?)
    RETURNING id, name, description, company_id, location, remote, contract_type, link, salary
`, opportunity.Name, opportunity.Description, opportunity.CompanyID, opportunity.Location, opportunity.Remote, opportunity.ContractType, opportunity.Link, opportunity.Salary).Scan(&newOpportunity).Error
	if err != nil {
			return models.Opportunity{}, err
	}

	for _, technology := range opportunity.Technologies {
		newOpportunityTechnology := models.NewOpportunityTechnology(newOpportunity.ID, technology)
		r.db.Exec(`
		CREATE TABLE IF NOT EXISTS opportunity_technologies (
				id INTEGER PRIMARY KEY,
				opportunity_id INTEGER,
				technology_id INTEGER,
				created_at TIMESTAMP,
				updated_at TIMESTAMP
		)
		`)
		err := r.db.Raw(`
		INSERT INTO opportunity_technologies (
				opportunity_id,
				technology_id
		) VALUES (?, ?)
		RETURNING id, opportunity_id, technology_id
		`, newOpportunityTechnology.OpportunityID, newOpportunityTechnology.TechnologyID).Scan(&newOpportunityTechnology).Error
		if err != nil {
			return models.Opportunity{}, err
		}
	}
	fmt.Println(newOpportunity)
	return newOpportunity, nil
}