package middleware

import (
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

func AuthMiddleware(app *fiber.App) {
	app.Use("/api", checkAuth)
}

func checkAuth(c *fiber.Ctx) error {
	jwtToken := c.Cookies("jwt")
	if jwtToken == "" {
		return c.Status(fiber.StatusUnauthorized).SendString("missing token")
	}

	claims := jwt.MapClaims{}
	jwtSecret := "secret"
	token, err := jwt.ParseWithClaims(jwtToken, claims, func(t *jwt.Token) (interface{}, error) {
		return []byte(jwtSecret), nil
	})

	if err != nil || !token.Valid {
		return c.Status(fiber.StatusUnauthorized).SendString("invalid token")
	}

	c.Locals("authClaims", claims)

	return c.Next()
}
