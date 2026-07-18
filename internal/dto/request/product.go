package request

type ProductRequest struct {
	ProductName string  `json:"productName" validate:"required"`
	Description string  `json:"description" validate:"required,min=10,max=100"`
	Price       float64 `json:"price" validate:"required"`
	Stock       int     `json:"stock" validate:"required"`
}

