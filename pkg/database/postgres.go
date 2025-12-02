package database

import (
	"log"
	"os"

	"service-nexa/internal/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func ConnectPostgres() (*gorm.DB, error) {
	var logMode logger.Interface
	if os.Getenv("ENV") == "prod" {
		logMode = logger.Default.LogMode(logger.Silent)
	} else {
		logMode = logger.Default.LogMode(logger.Info)
	}
	db, err := gorm.Open(postgres.Open(os.Getenv("DB_URL")), &gorm.Config{
		Logger: logMode,
	})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}
	log.Println("=== Connected to Database successfully! ===")
	db.AutoMigrate(&models.User{})
	return db, nil
}
