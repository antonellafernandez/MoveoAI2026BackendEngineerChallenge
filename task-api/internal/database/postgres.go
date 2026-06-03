package database

import (
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"task-api/internal/models"
)

func Connect() (*gorm.DB, error) {

	dsn := "host=db user=postgres password=postgres dbname=tasks port=5432 sslmode=disable"

	var db *gorm.DB
	var err error

	for i := 0; i < 10; i++ {

		db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

		if err == nil {
			break
		}

		time.Sleep(2 * time.Second)
	}

	if err != nil {
		return nil, err
	}

	db.AutoMigrate(&models.Task{})

	return db, nil
}
