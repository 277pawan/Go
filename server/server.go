package server

import (
	"go_ecommerce-app/configs"
	"go_ecommerce-app/internal/helper"
	"go_ecommerce-app/internal/routes"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func StartServer() {

	err := godotenv.Load()

	if err != nil {
		log.Println("Error loading .env file", err)
	}

	configs.ConnectDB()
	configs.ConnectRedis()

	// helper.KafkaProducer()
	go helper.KafkaConsumer()

	port := os.Getenv("PORT")

	app := fiber.New()

	routes.SetupRoutes(app)

	if err := app.Listen(":" + port); err != nil {
		log.Fatal("Server failed to start: ", err)
	}
}
