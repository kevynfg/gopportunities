package repositories

import (
	"time"

	"github.com/kevynfg/gopportunities/domain/models"
	"github.com/kevynfg/gopportunities/infra/logger"
	"gorm.io/gorm"
)

type TechnologiesRepositorySql struct {
	db *gorm.DB
}

func NewTechnologiesRepositorySql(db *gorm.DB) *TechnologiesRepositorySql {
	return &TechnologiesRepositorySql{db: db}
}

func (r *TechnologiesRepositorySql) CreateTechnology(technology *models.Technology) (models.Technology, error) {
	r.db.Exec("CREATE TABLE IF NOT EXISTS technologies (id INTEGER PRIMARY KEY, name TEXT, stack TEXT, created_at TIMESTAMP, updated_at TIMESTAMP)")
	var newTechnology models.Technology
	err := r.db.Raw("INSERT INTO technologies (name, stack, created_at) VALUES (?, ?, ?) RETURNING id, name, stack", technology.Name, technology.Stack, time.Now()).Scan(&newTechnology).Error
	if err != nil {
		logger.NewLogger("repositories/technology_repository.go").Errf("Error creating technology: %v", err)
		return models.Technology{}, err
	}
	return newTechnology, nil
}

func (r *TechnologiesRepositorySql) FindAll() ([]models.Technology, error) {
	var technologies []models.Technology
	err := r.db.Raw("SELECT * FROM technologies").Scan(&technologies).Error
	if err != nil {
		logger.NewLogger("repositories/technology_repository.go").Errf("Error finding technologies: %v", err)
		return nil, err
	}
	
	return technologies, nil
}