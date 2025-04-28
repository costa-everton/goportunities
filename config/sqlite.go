package config

import (
	"os"

	"github.com/costa-everton/goportunities/schemas"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func InitializeSQLite() (*gorm.DB, error){
	logger := GetLogger("sqlite")
	dbPath := "./db/main.db"

	// Check if database file exists
	_, err := os.Stat(dbPath)
	if os.IsNotExist(err){
		logger.Info("database file not found, creating...")

		// Create the database file and directory
		err = os.MkdirAll("./db", os.ModePerm)
		if err != nil {
			return nil, err
		}
		
		file, err := os.Create(dbPath)
		if err != nil {
			return nil, err
		}

		// Close the file
		file.Close()
	}

	// Create data base and connect
	db, err := gorm.Open(sqlite.Open(dbPath), &gorm.Config{})
	if err != nil{
		logger.Errorf("sqlite opening error: %v", err)
		return nil, err
	}

	// Migrate the schemas
	err = db.AutoMigrate(&schemas.Opportunities{})
	if err != nil {
		logger.Errorf("sqlite automigration error: %v", err)
		return nil, err
	}

	// Return DB
	return db, nil
}	