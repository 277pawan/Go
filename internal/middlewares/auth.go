package middlewares

import (
	"go_ecommerce-app/internal/helper"
	"strings"

	"github.com/gofiber/fiber/v2"
)

func Authenticate(c *fiber.Ctx) error {

	authHeader := c.Get("Authorization")

	if authHeader == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "Unauthorized",
		})
	}

	parts := strings.Split(authHeader, " ")

	if len(parts) != 2 || parts[0] != "Bearer" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "Invalid token",
		})
	}
	token := parts[1]

	claims, err := helper.VerifyJWT(token)

	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "Token Expired",
			"err":     err.Error(),
		})
	}

	c.Locals("user", claims)

	return c.Next()

}
