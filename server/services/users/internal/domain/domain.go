package domain

import (
	"gorm.io/gorm"

	"github.com/gofiber/fiber/v2"
)

type User struct {
	gorm.Model
	FullName string
	Email    string
	Password string
}

type Query struct {
	FullName *string `query:"full_name"`
	Email    *string `query:"email"`
	Password *string `query:"password"`
}

type UserRepository interface {
	FindAll() ([]*User, error)
	FindByID(id int64) (*User, error)
	Create(body *User) error
	Update(id int64, body *User) error
	Delete(id int64) error
	FindByFullName(filter *Query) ([]*User, error)
	FindByEmail(filter *Query) ([]*User, error)
	FindByPassword(filter *Query) ([]*User, error)
	FindBySearch(filter *Query) ([]*User, error)
}

type UserService interface {
	GetAllUser(filter *Query) ([]*User, error)
	GetUser(id int64) (*User, error)
	CreateUser(body *User) error
	UpdateUser(id int64, body *User) (*User, error)
	DeleteUser(id int64) error
}

type UserHandler interface {
	GetAllUser(c *fiber.Ctx) error
	GetUser(c *fiber.Ctx) error
	CreateUser(c *fiber.Ctx) error
	UpdateUser(c *fiber.Ctx) error
	DeleteUser(c *fiber.Ctx) error
}
