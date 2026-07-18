package container

import (
	"go_ecommerce-app/internal/handlers"
	"go_ecommerce-app/internal/repository"
	"go_ecommerce-app/internal/service"

	"gorm.io/gorm"
)

// Container holds all initialized handlers.
// Add new handlers here as the app grows.
type Container struct {
	UserHandler *handlers.UserHandler
	// ProductHandler *handlers.ProductHandler
	// OrderHandler   *handlers.OrderHandler
}

// New wires the full dependency graph and returns a ready Container.
func New(db *gorm.DB) *Container {
	// Repositories
	userRepo := repository.NewUserRepository(db)

	// Services
	userService := service.NewUserService(userRepo)

	// Handlers
	return &Container{
		UserHandler: handlers.NewUserHandler(userService),
	}
}
