package users

import "github.com/gofiber/fiber/v2"

func Router(app *fiber.App, h userHandler) {
	g := app.Group("/api/users")
	g.Get("", h.getAllUser)
	g.Get("/:id", h.getUser)
	g.Post("/", h.createUser)
	g.Put("/:id", h.updateUser)
	g.Delete("/:id", h.deleteUser)
}
