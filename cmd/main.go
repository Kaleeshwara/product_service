package main

import (
	"log"

	"product/handlers"
	"product/models"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	// Connect to PostgreSQL database
	dsn := "host=localhost user=postgres password=password dbname=product port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database: ", err)
	}

	// Auto Migrate the Products struct
	err = db.AutoMigrate(&models.Product{})
	if err != nil {
		log.Fatal("Failed to migrate database schema: ", err)
	}

	// Initialize Gin router
	router := gin.Default()

	// Inject the database instance into handlers
	handlers.SetupRoutes(router, db)

	// Run the server
	router.Run(":8080")
}
