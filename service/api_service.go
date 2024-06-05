package service

import (
	models "product/models"

	"gorm.io/gorm"
)

type APIService interface {
	SetDB(db *gorm.DB)
	GetProducts() ([]models.Product, error)
	GetProduct(id int) (models.Product, error)
	CreateProduct(product *models.Product) (*models.Product, error)
	UpdateProduct(id int, product *models.Product) (*models.Product, error)
	DeleteProduct(id int) error
}
