package routes

import (
	"go_ecommerce-app/internal/handlers"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {

	userHandler := &handlers.UserHandler{}
	UserRoutes(app, userHandler)
}
