package auth

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type Auth struct {
	gorm.Model
	Username string
	Password string
}

type authRepository interface {
	findByUsername(body *Auth) error
	findById(id int64) (*Auth, error)
	create(body *Auth) error
	update(body *Auth, id int64) error
	delete(id int64) error
}

type authService interface {
	signup(body *Auth) error
	signin(body *Auth) (string, error)
}

type authHandler interface {
	signup(c *fiber.Ctx) error
	signin(c *fiber.Ctx) error
	signout(c *fiber.Ctx) error
}
