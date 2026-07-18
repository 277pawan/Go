package handlers

import (
	"fmt"
	"go_ecommerce-app/internal/dto/request"
	"log"

	"github.com/gofiber/fiber/v2"
)

type UserHandler struct {
}

func (u *UserHandler) RegisterUser(c *fiber.Ctx) error {
	req := request.UserRequest{}

	fmt.Println(c)

	err := c.BodyParser(&req)
	if err != nil {
		log.Println("Error parsing request body:", err)
	}

	return nil
}
