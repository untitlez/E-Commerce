package users

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

type userRepository interface {
	findAll() ([]*User, error)
	findByID(id int64) (*User, error)
	create(body *User) error
	update(id int64, body *User) error
	delete(id int64) error
	findByFullName(filter *query) ([]*User, error)
	findByEmail(filter *query) ([]*User, error)
	findByPassword(filter *query) ([]*User, error)
	findBySearch(filter *query) ([]*User, error)
}

type userService interface {
	getAllUser(filter *query) ([]*User, error)
	getUser(id int64) (*User, error)
	createUser(body *User) error
	updateUser(id int64, body *User) (*User, error)
	deleteUser(id int64) error
}

type userHandler interface {
	getAllUser(c *fiber.Ctx) error
	getUser(c *fiber.Ctx) error
	createUser(c *fiber.Ctx) error
	updateUser(c *fiber.Ctx) error
	deleteUser(c *fiber.Ctx) error
}
