package service

import (
	"errors"
	"fmt"

	"go_ecommerce-app/internal/dto/request"
	"go_ecommerce-app/internal/dto/response"
	"go_ecommerce-app/internal/helper"
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

func (u *UserService) RegisterUser(req request.UserRequest) (res *response.UserResponse, err error) {

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

	return &response.UserResponse{
		Name: newUser.Name, Email: newUser.Email,
		CreatedAt: newUser.CreatedAt,
		UpdatedAt: newUser.UpdatedAt,
		DeletedAt: newUser.DeletedAt,
	}, nil

}

func (u *UserService) LoginUser(req *request.LoginRequest) (res *response.LoginResponse, err error) {
	user, err := u.repo.FindByEmail(req.Email)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("Invalid email and password.")
		}
		return nil, err
	}

	// NOTE: first arg = hashed (from DB), second arg = plain text (from request)
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password))

	if err != nil {
		return nil, errors.New("Invalid email and password.")
	}

	token, err := helper.GenerateJWT(user.Email, user.ID)
	if err != nil {
		return nil, err
	}

	return &response.LoginResponse{
		ID:        user.ID,
		Name:      user.Name,
		Email:     user.Email,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
		Token:     token,
	}, nil

}
