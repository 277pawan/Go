package repository

import (
	"go_ecommerce-app/internal/models"

	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{
		db: db,
	}
}

func (u *UserRepository) Create(user *models.User) error {
	return u.db.Create(user).Error
}

func (u *UserRepository) FindByEmail(email string) (*models.User, error) {
	var user models.User

	err := u.db.Where("email = ?", email).First(&user).Error
	if err != nil {
		return nil, err
	}

	return &user, nil
}
