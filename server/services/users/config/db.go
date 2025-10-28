package config

import (
	"log"
	"os"
	"server/services/users/internal/domain"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func InitDB() *gorm.DB {
	dsn := os.Getenv("DATABASE_URL")
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default,
	})

	if err != nil {
		log.Fatalf("failed to to connect to database: %v", err.Error())
	}

	if err = db.AutoMigrate(&domain.User{}); err != nil {
		log.Fatalf("failed to auto migrate Auth model: %v", err)
	}

	return db
}
