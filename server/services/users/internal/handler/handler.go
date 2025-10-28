package handler

import (
	"server/services/users/internal/domain"

	"github.com/gofiber/fiber/v2"
)

type handler struct {
	sv domain.UserService
}

type responseById struct {
	ID       uint   `json:"id"`
	FullName string `json:"full_name"`
	Email    string `json:"email"`
}

func NewHandler(s domain.UserService) *handler {
	return &handler{sv: s}
}

// Get All
func (h *handler) GetAllUser(c *fiber.Ctx) error {
	q := &domain.Query{}
	if errQueryParser := c.QueryParser(q); errQueryParser != nil {
		return errQueryParser
	}

	res, err := h.sv.GetAllUser(q)
	if err != nil {
		return c.Status(fiber.StatusNotFound).SendString(err.Error())
	}

	return c.Status(fiber.StatusOK).JSON(res)
}

// Get ID
func (h *handler) GetUser(c *fiber.Ctx) error {
	id, errParamsInt := c.ParamsInt("id")
	if errParamsInt != nil {
		return errParamsInt
	}

	res, err := h.sv.GetUser(int64(id))
	if err != nil {
		return c.Status(fiber.StatusNotFound).SendString(err.Error())
	}

	response := responseById{
		ID:       res.ID,
		FullName: res.FullName,
		Email:    res.Email,
	}

	return c.Status(fiber.StatusOK).JSON(response)
}

// Create
func (h *handler) CreateUser(c *fiber.Ctx) error {
	user := &domain.User{}
	if errBodyParser := c.BodyParser(user); errBodyParser != nil {
		return errBodyParser
	}

	if err := h.sv.CreateUser(user); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	return c.Status(fiber.StatusCreated).SendString("Create Success")
}

// Update
func (h *handler) UpdateUser(c *fiber.Ctx) error {
	user := &domain.User{}
	id, errParamsInt := c.ParamsInt("id")
	if errParamsInt != nil {
		return errParamsInt
	}

	if errBodyParser := c.BodyParser(user); errBodyParser != nil {
		return errBodyParser
	}

	_, err := h.sv.UpdateUser(int64(id), user)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	return c.Status(fiber.StatusOK).SendString("Update Success")
}

// Delete
func (h *handler) DeleteUser(c *fiber.Ctx) error {
	id, errParamsInt := c.ParamsInt("id")
	if errParamsInt != nil {
		return errParamsInt
	}

	if err := h.sv.DeleteUser(int64(id)); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	return c.Status(fiber.StatusOK).SendString("Delete Success")
}
