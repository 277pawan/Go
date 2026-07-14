package main

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
)

func main() {

	fmt.Println("hello")

	app := fiber.New()

	// Routes

	helperfunction()
	app.Listen("localhost:9000")
}

func helperfunction() {
	fmt.Println("hello this is a helper function")
}
