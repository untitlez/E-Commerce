package domain

import (
	"gorm.io/gorm"

	"github.com/gofiber/fiber/v2"
)

type Auth struct {
	gorm.Model
	Username string
	Password string
}

type AuthRepository interface {
	FindByUsername(body *Auth) (*Auth, error)
	FindById(id int64) (*Auth, error)
	Create(body *Auth) error
	Update(body *Auth, id int64) error
	Delete(id int64) error
}

type AuthService interface {
	SignUp(body *Auth) error
	SignIn(body *Auth) (string, error)
}

type AuthHandler interface {
	SignUp(c *fiber.Ctx) error
	SignIn(c *fiber.Ctx) error
	SignOut(c *fiber.Ctx) error
}
