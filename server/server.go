package server

import (
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"go_ecommerce-app/configs"
	"go_ecommerce-app/internal/routes"
	"log"
	"os"
)

func StartServer() {

	err := godotenv.Load()

	if err != nil {
		log.Println("Error loading .env file", err)
	}

	configs.ConnectDB()

	port := os.Getenv("PORT")

	app := fiber.New()

	routes.SetupRoutes(app)

	app.Listen(":" + port)
}

