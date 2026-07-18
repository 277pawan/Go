package models

import "gorm.io/gorm"

type Product struct {
	gorm.Model

	ProductName string
	Description string
	Price       float64
	Stock       int
}
