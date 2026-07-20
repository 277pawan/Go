package repository

import (
	"go_ecommerce-app/internal/models"

	"gorm.io/gorm"
)

type ProductRepository struct {
	db *gorm.DB
}

func NewProductRepository(db *gorm.DB) *ProductRepository {
	return &ProductRepository{
		db: db,
	}
}

func (p *ProductRepository) CreateProduct(product *models.Product) error {
	return p.db.Create(product).Error
}

func (p *ProductRepository) GetProducts() ([]models.Product, error) {
	var products []models.Product
	err := p.db.Preload("User").Find(&products).Error

	return products, err
}
