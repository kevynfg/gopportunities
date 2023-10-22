package repositories

import (
	"time"

	"github.com/kevynfg/gopportunities/domain/models"
	"github.com/kevynfg/gopportunities/infra/logger"
	"gorm.io/gorm"
)

type CompaniesRepositorySql struct {
	db *gorm.DB
}

func NewCompaniesRepositorySql(db *gorm.DB) *CompaniesRepositorySql {
	return &CompaniesRepositorySql{db: db}
}

func (r *CompaniesRepositorySql) CreateCompany(company *models.Company) (models.Company, error) {
	r.db.Exec("CREATE TABLE IF NOT EXISTS company (id INTEGER PRIMARY KEY, name TEXT, startup BOOL, created_at TIMESTAMP, updated_at TIMESTAMP)")
	var newCompany models.Company
	err := r.db.Raw("INSERT INTO company (name, startup, created_at) VALUES (?, ?, ?) RETURNING id, name, startup, created_at", company.Name, company.Startup, time.Now()).Scan(&newCompany).Error
	if err != nil {
		logger.NewLogger("repositories/company_repository.go").Errf("Error creating company: %v", err)
		return models.Company{}, err
	}
	return newCompany, nil
}

func (r *CompaniesRepositorySql) FindAll() ([]*models.Company, error) {
	var companies []*models.Company
	err := r.db.Raw("SELECT * FROM company").Scan(&companies).Error
	if err != nil {
		logger.NewLogger("repositories/company_repository.go").Errf("Error finding companies: %v", err)
		return nil, err
	}
	
	return companies, nil
}