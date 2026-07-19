package handlers

import (
	"go_ecommerce-app/internal/dto/request"
	"go_ecommerce-app/internal/service"
	"go_ecommerce-app/internal/validation"

	"github.com/gofiber/fiber/v2"
)

type UserHandler struct {
	userService *service.UserService
}

func NewUserHandler(service *service.UserService) *UserHandler {
	return &UserHandler{
		userService: service,
	}
}

// Register User Hanlder
func (u *UserHandler) RegisterUser(c *fiber.Ctx) error {

	var req request.UserRequest

	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid request",
		})
	}

	if err := validation.Validate.Struct(req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"errors": validation.FormatValidationErrors(err),
		})
	}

	newUser, err := u.userService.RegisterUser(req)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{"data": newUser,
		"message": "User registered successfully",
	})
}

// Login User Handler
func (u *UserHandler) LoginUser(c *fiber.Ctx) error {

	var req request.LoginRequest

	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid request",
		})
	}

	if err := validation.Validate.Struct(req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"errors": validation.FormatValidationErrors(err),
		})
	}

	userLogin, err := u.userService.LoginUser(&req)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return c.Status(fiber.StatusAccepted).JSON(fiber.Map{
		"data":    userLogin,
		"message": "User Login Successfully",
	})
}
