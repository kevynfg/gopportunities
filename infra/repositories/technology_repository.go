package repositories

import (
	"github.com/kevynfg/gopportunities/domain/models"
	"gorm.io/gorm"
)

type TechnologiesRepositorySql struct {
	db *gorm.DB
}

func NewTechnologiesRepositorySql(db *gorm.DB) *TechnologiesRepositorySql {
	return &TechnologiesRepositorySql{db: db}
}

func (r *TechnologiesRepositorySql) CreateTechnology(technology *models.Technology) (models.Technology, error) {
	r.db.Exec("CREATE TABLE IF NOT EXISTS technologies (id INTEGER PRIMARY KEY, name TEXT)")
	var newTechnology models.Technology
	err := r.db.Raw("INSERT INTO technologies (name) VALUES (?) RETURNING id, name", technology.Name).Scan(&newTechnology).Error
	if err != nil {
		return models.Technology{}, err
	}
	return newTechnology, nil
}

func (r *TechnologiesRepositorySql) FindAll() ([]models.Technology, error) {
	var technologies []models.Technology
	err := r.db.Raw("SELECT * FROM technologies").Scan(&technologies).Error
	if err != nil {
		return nil, err
	}
	
	return technologies, nil
}