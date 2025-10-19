package auth

import (
	"time"

	"github.com/gofiber/fiber/v2"
)

type handler struct {
	sv authService
}

type responseAuth struct {
	Id       uint   `json:"id"`
	Username string `json:"username"`
}

func NewHandler(sv authService) *handler {
	return &handler{sv: sv}
}

func (h *handler) signup(c *fiber.Ctx) error {
	return c.JSON("")
}

func (h *handler) signin(c *fiber.Ctx) error {
	body := &Auth{}
	if errBodyParser := c.BodyParser(body); errBodyParser != nil {
		return errBodyParser
	}

	signedToken, err := h.sv.signin(body)
	if err != nil {
		return err
	}

	c.Cookie(&fiber.Cookie{
		Name:     "jwt",
		Value:    signedToken,
		Expires:  time.Now().Add(time.Hour),
		Secure:   true,
		HTTPOnly: true,
		SameSite: "Strict",
	})

	return c.Status(fiber.StatusOK).SendString("signin success")
}

func (h *handler) signout(c *fiber.Ctx) error {
	c.Cookie(&fiber.Cookie{
		Name:     "jwt",
		Value:    "",
		Expires:  time.Now().Add(-time.Hour),
		HTTPOnly: true,
	})

	return c.Status(fiber.StatusOK).SendString("signout success")
}
