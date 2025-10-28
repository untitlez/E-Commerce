package middleware

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func CORS(app *fiber.App) {
	app.Use(cors.New(cors.Config{
		AllowOrigins:     "https://e-commerce-iota-six-84.vercel.app/",
		AllowCredentials: true,
	}))
}
