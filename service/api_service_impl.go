package service

import (
	models "product/models"

	"gorm.io/gorm"
)

type ProductHandler struct {
	db *gorm.DB
}

func (productsHandler *ProductHandler) SetDB(db *gorm.DB) {
	productsHandler.db = db
}

func (productsHandler *ProductHandler) GetProducts() ([]models.Product, error) {
	var products []models.Product
	result := productsHandler.db.Find(&products)
	return products, result.Error
}

func (productsHandler *ProductHandler) GetProduct(id int) (models.Product, error) {
	var product models.Product
	result := productsHandler.db.First(&product, id)
	return product, result.Error
}

func (productsHandler *ProductHandler) CreateProduct(product *models.Product) (*models.Product, error) {
	result := productsHandler.db.Create(&product)
	return product, result.Error
}

func (productsHandler *ProductHandler) UpdateProduct(id int, product *models.Product) (*models.Product, error) {
	var existingProduct models.Product
	if err := productsHandler.db.First(&existingProduct, id).Error; err != nil {
		return nil, err
	}
	if err := productsHandler.db.Model(&existingProduct).Updates(product).Error; err != nil {
		return nil, err
	}
	return product, nil
}

func (productsHandler *ProductHandler) DeleteProduct(id int) error {
	var product models.Product
	if err := productsHandler.db.First(&product, id).Error; err != nil {
		return err
	}
	if err := productsHandler.db.Delete(&product).Error; err != nil {
		return err
	}
	return nil
}

func NewProductService(db *gorm.DB) APIService {
	return &ProductHandler{db: db}
}
