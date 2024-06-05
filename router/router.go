package router

import (
	"product/handlers"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"gorm.io/gorm"
)

func SetupRoutes(router *gin.Engine, db *gorm.DB) *gin.Engine {
	productHandler := handlers.NewProductHandler(db)
	router.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	product := router.Group("/products")
	{
		product.GET("/", productHandler.GetProducts)
		product.GET("/:id", productHandler.GetProduct)
		product.POST("/", productHandler.CreateProduct)
		product.PUT("/:id", productHandler.UpdateProduct)
		product.DELETE("/:id", productHandler.DeleteProduct)
	}
	return router
}
