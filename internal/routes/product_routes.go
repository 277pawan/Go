package routes

import (
	"go_ecommerce-app/internal/handlers"

	"github.com/gofiber/fiber/v2"
)

func ProductRoutes(app *fiber.App, handler *handlers.ProductHandler) {
	app.Post("/product", handler.CreateProduct)
}

