package main

import (
	"server/services/users/config"
	"server/services/users/internal/handler"
	"server/services/users/internal/repository"
	"server/services/users/internal/service"

	"github.com/gofiber/fiber/v2"
)

func main() {
	db := config.InitDB()

	usersRepository := repository.NewRepository(db)
	usersService := service.NewService(usersRepository)
	usersHandler := handler.NewHandler(usersService)

	app := fiber.New()

	app.Get("/api/users", usersHandler.GetAllUser)
	app.Get("/api/users/:id", usersHandler.GetUser)
	app.Post("/api/users", usersHandler.CreateUser)
	app.Put("/api/users/:id", usersHandler.UpdateUser)
	app.Delete("/api/users/:id", usersHandler.DeleteUser)

	app.Listen(":5000")
}
