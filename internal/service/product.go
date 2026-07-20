package service

import (
	"go_ecommerce-app/internal/dto/request"
	"go_ecommerce-app/internal/dto/response"
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
		UserID:      req.UserID,
	}

	err = p.repo.CreateProduct(newProduct)
	if err != nil {
		return nil, err
	}

	return newProduct, err
}

// Get Products API
func (p *ProdcutService) GetProducts() ([]response.ProductResponse, error) {

	var responses []response.ProductResponse

	products, err := p.repo.GetProducts()
	if err != nil {
		return nil, err
	}

	for _, product := range products {
		responses = append(responses, response.ProductResponse{
			ProductName: product.ProductName,
			Description: product.Description,
			Price:       product.Price,
			Stock:       product.Stock,
			User: response.ProductUserResponse{
				ID:    product.User.ID,
				Name:  product.User.Name,
				Email: product.User.Email,
			},
		})
	}

	return responses, err

}
