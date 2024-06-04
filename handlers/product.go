package handlers

import (
	"net/http"
	"strconv"

	"product/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupRoutes(router *gin.Engine, db *gorm.DB) {
	// Initialize the products handler with the database instance
	productsHandler := ProductHandler{db}

	router.GET("/products", productsHandler.getProducts)
	router.GET("/products/:id", productsHandler.getProduct)
	router.POST("/products", productsHandler.createProduct)
	router.PUT("/products/:id", productsHandler.updateProduct)
	router.DELETE("/products/:id", productsHandler.deleteProduct)
}

// ProductHandler handles product CRUD operations
type ProductHandler struct {
	db *gorm.DB
}

// @Summary Get all products
// @Description Get all products
// @Tags Products
// @Produce json
// @Success 200 {array} Product
// @Router /products [get]
func (ph *ProductHandler) getProducts(c *gin.Context) {
	var products []models.Product
	ph.db.Find(&products)
	c.JSON(http.StatusOK, products)
}

// @Summary Get a product by ID
// @Description Get a product by ID
// @Tags Products
// @Produce json
// @Param id path int true "Product ID"
// @Success 200 {object} Product
// @Failure 404 {object} ErrorResponse
// @Router /products/{id} [get]
func (ph *ProductHandler) getProduct(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var product models.Product
	if err := ph.db.First(&product, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
		return
	}
	c.JSON(http.StatusOK, product)
}

// @Summary Create a new product
// @Description Create a new product
// @Tags Products
// @Accept json
// @Produce json
// @Param product body Product true "Product object"
// @Success 201 {object} Product
// @Failure 400 {object} ErrorResponse
// @Router /products [post]
func (ph *ProductHandler) createProduct(c *gin.Context) {
	var newProduct models.Product
	if err := c.BindJSON(&newProduct); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ph.db.Create(&newProduct)
	c.JSON(http.StatusCreated, newProduct)
}

// @Summary Update a product by ID
// @Description Update a product by ID
// @Tags Products
// @Accept json
// @Produce json
// @Param id path int true "Product ID"
// @Param product body Product true "Product object"
// @Success 200 {object} Product
// @Failure 400 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
// @Router /products/{id} [put]
func (ph *ProductHandler) updateProduct(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var updatedProduct models.Product
	if err := c.BindJSON(&updatedProduct); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := ph.db.First(&models.Product{}, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
		return
	}
	ph.db.Model(&models.Product{}).Where("id = ?", id).Updates(updatedProduct)
	c.JSON(http.StatusOK, updatedProduct)
}

// @Summary Delete a product by ID
// @Description Delete a product by ID
// @Tags Products
// @Param id path int true "Product ID"
// @Success 200 {object} SuccessResponse
// @Failure 404 {object} ErrorResponse
// @Router /products/{id} [delete]
func (ph *ProductHandler) deleteProduct(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var product models.Product
	if err := ph.db.First(&product, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
		return
	}
	ph.db.Delete(&product)
	c.JSON(http.StatusOK, gin.H{"message": "Product deleted"})
}
