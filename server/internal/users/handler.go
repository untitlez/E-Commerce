package users

import "github.com/gofiber/fiber/v2"

// "time"

type handler struct {
	sv userService
}

type responseById struct {
	ID       uint   `json:"id"`
	FullName string `json:"full_name"`
	Email    string `json:"email"`
}

type query struct {
	FullName *string `query:"full_name"`
	Email    *string `query:"email"`
	Password *string `query:"password"`
}

func NewHandler(s userService) *handler {
	return &handler{sv: s}
}

// Get All
func (h *handler) getAllUser(c *fiber.Ctx) error {
	q := &query{}
	if errQueryParser := c.QueryParser(q); errQueryParser != nil {
		return errQueryParser
	}

	res, err := h.sv.getAllUser(q)
	if err != nil {
		return c.Status(fiber.StatusNotFound).SendString(err.Error())
	}

	return c.Status(fiber.StatusOK).JSON(res)
}

// Get ID
func (h *handler) getUser(c *fiber.Ctx) error {
	id, errParamsInt := c.ParamsInt("id")
	if errParamsInt != nil {
		return errParamsInt
	}

	res, err := h.sv.getUser(int64(id))
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
func (h *handler) createUser(c *fiber.Ctx) error {
	user := &User{}
	if errBodyParser := c.BodyParser(user); errBodyParser != nil {
		return errBodyParser
	}

	if err := h.sv.createUser(user); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	return c.Status(fiber.StatusCreated).SendString("Create Success")
}

// Update
func (h *handler) updateUser(c *fiber.Ctx) error {
	user := &User{}
	id, errParamsInt := c.ParamsInt("id")
	if errParamsInt != nil {
		return errParamsInt
	}

	if errBodyParser := c.BodyParser(user); errBodyParser != nil {
		return errBodyParser
	}

	_, err := h.sv.updateUser(int64(id), user)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	return c.Status(fiber.StatusOK).SendString("Update Success")
}

// Delete
func (h *handler) deleteUser(c *fiber.Ctx) error {
	id, errParamsInt := c.ParamsInt("id")
	if errParamsInt != nil {
		return errParamsInt
	}

	if err := h.sv.deleteUser(int64(id)); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	return c.Status(fiber.StatusOK).SendString("Delete Success")
}
