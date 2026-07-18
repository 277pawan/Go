package service

import (
	"go_ecommerce-app/internal/dto/request"
	"go_ecommerce-app/internal/models"
	"go_ecommerce-app/internal/repository"
)

type ProdcutService struct {
	repo *repository.ProductRepository
}

func NewProductService(repo *repository.ProductRepository) *ProdcutService {
	return &ProdcutService{
		repo: repo,
	}
}

func (p *ProdcutService) CreateProduct(req request.ProductRequest) (res *models.Product, err error) {

	newProduct := &models.Product{
		ProductName: req.ProductName,
		Description: req.Description,
		Price:       req.Price,
		Stock:       req.Stock,
	}

	p.repo.CreateProduct(newProduct)

	return newProduct, err
}
