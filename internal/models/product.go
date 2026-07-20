package models

import "gorm.io/gorm"

type Product struct {
	gorm.Model

	ProductName string
	Description string
	Price       float64
	Stock       int

	UserID uint
	User   User
}

// if we wanna make email as a foreign key default it takes id as it.
//
//UserEmail string it will take it default UserEmail
//User User `gorm:"foreignKey:UserEmail;references:Email"`
//
