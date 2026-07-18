package configs

import (
	"go_ecommerce-app/internal/models"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {

	dsn := "host=localhost user=postgres password=root dbname=golearn port=5432"

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}
	db.AutoMigrate(models.User{}, models.Product{})
	DB = db
	log.Println("Database connected and migrated successfully")
}
