package response

type PrdouctResponse struct {
	ProductName string `json:"product_name"`
	Description string `json:"product_description"`
	Price       float64
	Stock       int
}

