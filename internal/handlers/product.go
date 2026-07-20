// Package handlers for Product Handler
package handlers

import (
	"go_ecommerce-app/internal/dto/request"
	"go_ecommerce-app/internal/dto/response"
	"go_ecommerce-app/internal/service"
	"go_ecommerce-app/internal/validation"

	"github.com/gofiber/fiber/v2"
)

type ProductHandler struct {
	productService *service.ProdcutService
}

func NewProductHandler(service *service.ProdcutService) *ProductHandler {
	return &ProductHandler{
		productService: service,
	}
}

// Create Product API
func (p *ProductHandler) CreateProduct(c *fiber.Ctx) error {
	var req request.ProductRequest

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

	newProduct, err := p.productService.CreateProduct(req)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	productResponse := response.ProductCreateResponse{
		ProductName: newProduct.ProductName,
		Description: newProduct.Description,
		Price:       newProduct.Price,
		Stock:       newProduct.Stock,
		UserID:      newProduct.UserID,
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{"data": productResponse,
		"message": "Product Created successfully",
	})
}

func (p *ProductHandler) GetProducts(c *fiber.Ctx) error {

	products, err := p.productService.GetProducts()
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"data":    products,
		"message": "Products data fetched successfully.",
	})
}
