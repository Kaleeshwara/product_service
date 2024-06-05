package handlers

import (
	"net/http"
	"strconv"

	"product/service"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"product/models"
)

type ProductResponse struct {
	ID    uint    `json:"id"`
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}

type ProductRequest struct {
	ID    uint    `json:"id"`
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}

type Response map[string]interface{}

type ProductHandler struct {
	service service.APIService
}

func NewProductHandler(db *gorm.DB) *ProductHandler {
	return &ProductHandler{service: service.NewProductService(db)}
}

// @Summary Get all products
// @Description Get all products
// @Tags Products
// @Produce json
// @Success 200
// @Router /products [get]
func (h *ProductHandler) GetProducts(c *gin.Context) {
	products, err := h.service.GetProducts()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get products"})
		return
	}
	c.JSON(http.StatusOK, products)
}

// @Summary Get a product by ID
// @Description Get a product by its ID
// @Tags Products
// @Accept json
// @Produce json
// @Param id path int true "Product ID"
// @Success 200 {object} ProductResponse
// @Router /products/{id} [get]
func (h *ProductHandler) GetProduct(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	product, err := h.service.GetProduct(id)
	if err != nil {
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
// @Param product body ProductRequest true "Product details"
// @Success 201 {object} ProductResponse
// @Router /products [post]
func (h *ProductHandler) CreateProduct(c *gin.Context) {
	var newProduct models.Product
	if err := c.BindJSON(&newProduct); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	createdProduct, err := h.service.CreateProduct(&newProduct)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create product"})
		return
	}
	newProduct = *createdProduct
	c.JSON(http.StatusCreated, newProduct)
}

// @Summary Update an existing product
// @Description Update an existing product
// @Tags Products
// @Accept json
// @Produce json
// @Param id path int true "Product ID"
// @Param product body ProductRequest true "Updated product details"
// @Success 200 {object} ProductResponse
// @Router /products/{id} [put]
func (h *ProductHandler) UpdateProduct(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var updatedProduct models.Product
	if err := c.BindJSON(&updatedProduct); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	updatedProduct.ID = uint(id)
	result, err := h.service.UpdateProduct(id, &updatedProduct)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
		return
	}
	c.JSON(http.StatusOK, result)
}

// @Summary Delete a product
// @Description Delete a product
// @Tags Products
// @Param id path int true "Product ID"
// @Produce json
// @Success 200 {object} Response
// @Router /products/{id} [delete]
func (h *ProductHandler) DeleteProduct(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	err := h.service.DeleteProduct(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Product deleted"})
}
