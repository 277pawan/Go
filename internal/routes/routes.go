package routes

import (
	"go_ecommerce-app/configs"
	"go_ecommerce-app/internal/container"
	"go_ecommerce-app/internal/middlewares"
	"time"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	c := container.New(configs.DB)

	app.Use(middlewares.RateLimiter(configs.RedisDB, 5, 10*time.Second))

	UserRoutes(app, c.UserHandler)
	ProductRoutes(app, c.ProductHandler)

	// ProductRoutes(app, c.ProductHandler)
	// OrderRoutes(app, c.OrderHandler)
}
