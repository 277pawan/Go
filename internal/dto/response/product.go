package response

type ProductCreateResponse struct {
	ProductName string `json:"product_name"`
	Description string `json:"product_description"`
	Price       float64
	Stock       int
	UserID      uint
}

type ProductUserResponse struct {
	ID    uint   `json:"user_id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

type ProductResponse struct {
	ProductName string `json:"product_name"`
	Description string `json:"product_description"`
	Price       float64
	Stock       int
	User        ProductUserResponse
}
