package config

import (
	"fmt"
	"os"

	"github.com/bagusrexy/test-dataon/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func CreateConnectionPostgres() (*gorm.DB, error) {
	dbUsername := os.Getenv("DATABASE_USERNAME")
	dbPassword := os.Getenv("DATABASE_PASSWORD")
	dbName := os.Getenv("DATABASE_NAME")
	dbHost := os.Getenv("DATABASE_HOST")
	dbPort := os.Getenv("DATABASE_PORT")
	dbSSLMode := os.Getenv("DATABASE_SSLMODE")

	dbConn := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s", dbHost, dbPort, dbUsername, dbName, dbPassword, dbSSLMode)

	db, err := gorm.Open(postgres.Open(dbConn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	err = models.RunMigrate(db)
	if err != nil {
		return nil, err
	}
	return db, nil
}
