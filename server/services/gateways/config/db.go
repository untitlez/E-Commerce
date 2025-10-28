package config

import (
	"fmt"
	"log"
	"server/services/gateways/internal/domain"

	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func IntiDB() *gorm.DB {
	dsn := fmt.Sprintf("postgresql://%v:%v@%v:%v/%v?sslmode=disable",
		viper.GetString("db.username"),
		viper.GetString("db.password"),
		viper.GetString("db.host"),
		viper.GetInt("db.port"),
		viper.GetString("db.database"),
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default,
	})

	if err != nil {
		log.Fatalf("failed to to connect to database: %v", err.Error())
	}

	if err = db.AutoMigrate(&domain.Auth{}); err != nil {
		log.Fatalf("failed to auto migrate Auth model: %v", err)
	}

	return db
}
