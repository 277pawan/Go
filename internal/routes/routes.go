package routes

import (
	"go_ecommerce-app/configs"
	"go_ecommerce-app/internal/container"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	c := container.New(configs.DB)

	UserRoutes(app, c.UserHandler)
	// ProductRoutes(app, c.ProductHandler)
	// OrderRoutes(app, c.OrderHandler)
}

