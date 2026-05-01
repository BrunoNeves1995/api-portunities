package config

import (
	"os"

	"github.com/BrunoNeves1995/api-portunities/schemas"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func InitializeSqlite() (*gorm.DB, error) {
	logger := GetLogger("sqlite")
	dbPath := "./db/portunings.db"

	//Check if the database exists
	_, err := os.Stat(dbPath)
	if os.IsNotExist(err) {
		logger.Errorf("database file not found, creating new database...")
		//Create the database file and directory
		err = os.MkdirAll("./db", os.ModePerm)
		if err != nil {
			logger.Errorf("error creating directory for database: %v", err)
			return nil, err
		}
		file, err := os.Create(dbPath)
		if err != nil {
			logger.Errorf("error creating file for database: %v", err)
			return nil, err
		}
		logger.Infof("database file created: %v", dbPath)
		file.Close()

	}

	//Create DB and Conect
	db, err := gorm.Open(sqlite.Open(dbPath), &gorm.Config{})
	if err != nil {
		logger.Errorf("sqlite opening error: %v", err)
		return nil, err
	}
	//Migrate the Schema
	err = db.AutoMigrate(&schemas.Opening{})
	if err != nil {
		logger.Errorf("sqlite migration error: %v", err)
		return nil, err
	}
	//Retorn the DB
	return db, nil
}
