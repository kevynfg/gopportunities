package db

import (
	"os"

	"github.com/kevynfg/gopportunities/domain/models"
	"github.com/kevynfg/gopportunities/infra/logger"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func InitDB() *gorm.DB {
	dbPath := "./infra/db/opportunities.db"
	logger := logger.GetLogger("db");
	
	_, err := os.Stat(dbPath); 
	if err == nil {
		logger.Infof("database file exists: %v", dbPath)
	} else if os.IsNotExist(err) {
		logger.Warningf("database file does not exist, creating it: %v", err)
		err = os.MkdirAll("./infra/db", os.ModePerm)
		if err != nil {
			logger.Errf("failed to create schema directory: %v", err)
			panic("failed to create schema directory")
		}
		file, err := os.Create(dbPath)
		if err != nil {
			logger.Errf("failed to create database file: %v", err)
			panic("failed to create database file")
		}
		file.Close()
	}

	db, err := gorm.Open(sqlite.Open(dbPath), &gorm.Config{})
	
	if err != nil {
		logger.Errf("failed to connect database: %v", err)
		panic("failed to connect database")
	}
	
	err = db.AutoMigrate(&models.Opportunity{})
	if err != nil {
		logger.Errf("failed to migrate database: %v", err)
		panic("failed to migrate database")
	}
	return db
}