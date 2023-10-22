package db

import (
	"log"
	"os"
	"time"

	"github.com/kevynfg/gopportunities/domain/models"
	customLogger "github.com/kevynfg/gopportunities/infra/logger"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var db *gorm.DB

func InitDB() *gorm.DB {
	dbPath := "./infra/db/opportunities.db"
	customLogger := customLogger.GetLogger("db");
	
	_, err := os.Stat(dbPath); 
	if err == nil {
		customLogger.Infof("database file exists: %v", dbPath)
	} else if os.IsNotExist(err) {
		customLogger.Warningf("database file does not exist, creating it: %v", err)
		err = os.MkdirAll("./infra/db", os.ModePerm)
		if err != nil {
			customLogger.Errf("failed to create schema directory: %v", err)
			panic("failed to create schema directory")
		}
		file, err := os.Create(dbPath)
		if err != nil {
			customLogger.Errf("failed to create database file: %v", err)
			panic("failed to create database file")
		}
		file.Close()
	}

	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold:              time.Second,   // Slow SQL threshold
			LogLevel:                   logger.Info, // Log level
			IgnoreRecordNotFoundError: true,           // Ignore ErrRecordNotFound error for logger
			ParameterizedQueries:      true,           // Don't include params in the SQL log
			Colorful:                  false,          // Disable color
		},
	)

	db, err := gorm.Open(sqlite.Open(dbPath), &gorm.Config{
		Logger: newLogger,
	})

	if err != nil {
		customLogger.Errf("failed to connect database: %v", err)
		panic("failed to connect database")
	}
	
	err = db.AutoMigrate(&models.Opportunity{}, &models.Company{}, &models.Technology{})
	if err != nil {
		customLogger.Errf("failed to migrate database: %v", err)
		panic("failed to migrate database")
	}
	return db
}

func GetDB() *gorm.DB {
	if db == nil {
		db = InitDB()
	}
	return db
}