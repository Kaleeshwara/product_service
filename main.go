package main

import (
	"fmt"
	"log"

	_ "product/docs"
	"product/models"

	config "product/config"
	routers "product/router"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// @title           Product Service
// @version         1.0
// @description     A CRUD API for the Product Service.
func main() {
	cfg := config.LoadEnv()

	dsn := fmt.Sprintf("host=%s port=%d user=%s dbname=%s password=%s sslmode=disable",
		cfg.Postgres.Host, cfg.Postgres.Port, cfg.Postgres.User, cfg.Postgres.Database, cfg.Postgres.Password)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("Failed to connect to database: ", err)
	}

	err = db.AutoMigrate(&models.Product{})

	if err != nil {
		log.Fatal("Failed to migrate database schema: ", err)
	}

	router := gin.Default()

	routers.SetupRoutes(router, db)

	router.Run(":8089")
}
