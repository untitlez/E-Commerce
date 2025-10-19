package auth

import "github.com/gofiber/fiber/v2"

func Router(app *fiber.App, h authHandler) {
	app.Group("/api/auth")
	// g.Get("", h)
	// g.Get("/:id", h)
	// g.Post("/", h)
	// g.Put("/:id", h)
	// g.Delete("/:id", h)
}
