package service

import (
	"errors"
	"fmt"

	"go_ecommerce-app/internal/dto/request"
	"go_ecommerce-app/internal/models"
	"go_ecommerce-app/internal/repository"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type UserService struct {
	repo *repository.UserRepository
}

func NewUserService(repo *repository.UserRepository) *UserService {
	return &UserService{
		repo: repo,
	}
}

func (u *UserService) RegisterUser(req request.UserRequest) (res *models.User, err error) {

	user, err := u.repo.FindByEmail(req.Email)

	if err == nil && user != nil {
		return nil, fmt.Errorf("user already exists")
	}

	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}

	hashPassword, err := bcrypt.GenerateFromPassword(
		[]byte(req.Password),
		bcrypt.DefaultCost,
	)

	if err != nil {
		return nil, err
	}

	newUser := &models.User{
		Name:     req.Name,
		Email:    req.Email,
		Password: string(hashPassword),
	}
	u.repo.Create(newUser)

	return newUser, nil
}
