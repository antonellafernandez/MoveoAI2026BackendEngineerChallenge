package database

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"task-api/internal/models"
)

func Connect() (*gorm.DB, error) {

	dsn := "host=db user=postgres password=postgres dbname=tasks port=5432 sslmode=disable"

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	if err := db.AutoMigrate(&models.Task{}); err != nil {
		return nil, err
	}

	return db, nil
}
