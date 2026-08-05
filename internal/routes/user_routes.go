package routes

import (
	"go_ecommerce-app/internal/handlers"
	"go_ecommerce-app/internal/middlewares"

	"github.com/gofiber/fiber/v2"
)

func UserRoutes(app *fiber.App, handler *handlers.UserHandler) {
	app.Post("/user", handler.RegisterUser)
	app.Post("/login", handler.LoginUser)
	app.Get("/users", middlewares.Authenticate, handler.GetUser)
}
