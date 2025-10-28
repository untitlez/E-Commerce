package main

import (
	"fmt"
	"server/services/gateways/config"
	"server/services/gateways/internal/handler"
	"server/services/gateways/internal/middleware"
	"server/services/gateways/internal/repository"
	"server/services/gateways/internal/service"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/proxy"
	"github.com/spf13/viper"
)

func main() {
	config.InitConfig()
	config.InitTimeZone()
	db := config.IntiDB()

	authRepository := repository.NewRepository(db)
	authService := service.NewService(authRepository)
	authHandler := handler.NewHandler(authService)

	app := fiber.New()
	middleware.CORS(app)
	middleware.AuthMiddleware(app)

	app.Post("/auth/signup", authHandler.SignUp)
	app.Post("/auth/signin", authHandler.SignIn)
	app.Post("/auth/signout", authHandler.SignOut)

	app.All("/api/users/*", proxy.Forward("http://user-services:8001"))

	v := viper.GetInt("app.port")
	fmt.Printf("Server start at port %v", v)
	app.Listen(fmt.Sprintf(":%v", v))
}
