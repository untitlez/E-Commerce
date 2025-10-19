package main

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"

	"server/internal/auth"
	"server/internal/users"

	"github.com/gofiber/fiber/v2"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "myuser"
	password = "mypassword"
	dbname   = "mydatabase"
)

func main() {
	dsn := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default,
	})

	if err != nil {
		panic("failed to connect to database")
	}

	db.AutoMigrate(users.User{}, auth.Auth{})

	app := fiber.New()

	uRepository := users.NewRepository(db)
	uService := users.NewService(uRepository)
	uHandler := users.NewHandler(uService)
	users.Router(app, uHandler)

	authRepository := auth.NewRepository(db)
	authService := auth.NewService(authRepository)
	authHandler := auth.NewHandler(authService)
	auth.Router(app, authHandler)

	app.Listen(":5000")
}
