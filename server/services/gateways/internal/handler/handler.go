package handler

import (
	"server/services/gateways/internal/domain"
	"time"

	"github.com/gofiber/fiber/v2"
)

type handler struct {
	sv domain.AuthService
}

func NewHandler(sv domain.AuthService) *handler {
	return &handler{sv: sv}
}

// SIGN UP
func (h *handler) SignUp(c *fiber.Ctx) error {
	body := &domain.Auth{}
	if errBodyParser := c.BodyParser(body); errBodyParser != nil {
		return errBodyParser
	}

	if err := h.sv.SignUp(body); err != nil {
		return err
	}

	return c.Status(fiber.StatusOK).SendString("sign up success")
}

// SIGN IN
func (h *handler) SignIn(c *fiber.Ctx) error {
	body := &domain.Auth{}
	if errBodyParser := c.BodyParser(body); errBodyParser != nil {
		return errBodyParser
	}

	signedToken, err := h.sv.SignIn(body)
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

	// อย่าลืมแก้ก่อน deploy
	// c.Status(fiber.StatusOK).SendString("sign in success")
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "sign in success",
		"token":   signedToken,
	})
}

// SIGN OUT
func (h *handler) SignOut(c *fiber.Ctx) error {
	c.Cookie(&fiber.Cookie{
		Name:    "jwt",
		Value:   "",
		Expires: time.Now().Add(-time.Hour),
	})

	if cookie := c.Cookies("jwt"); cookie == "" {
		return c.Status(fiber.StatusInternalServerError).SendString("failed to clear session cookie.")
	}

	return c.Status(fiber.StatusOK).SendString("sign out success")
}
