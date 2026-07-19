package routes

import (
	"go_ecommerce-app/internal/handlers"

	"github.com/gofiber/fiber/v2"
)

func UserRoutes(app *fiber.App, handler *handlers.UserHandler) {
	app.Post("/user", handler.RegisterUser)
	app.Post("/login", handler.LoginUser)
}
