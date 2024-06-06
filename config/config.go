package config

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	Postgres struct {
		Host     string
		Port     int
		User     string
		Password string
		Database string
		Schema   string
	}
}

// LoadEnv function to load environment variables from .env file
func LoadEnv() *Config {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	portStr := os.Getenv("DB_PORT")
	port, err := strconv.Atoi(portStr)
	if err != nil {
		log.Fatalf("Error converting DB_PORT to integer: %v", err)
	}

	fmt.Println(os.Getenv("DB_HOST"), os.Getenv("DB_PASSWORD"))
	return &Config{
		Postgres: struct {
			Host     string
			Port     int
			User     string
			Password string
			Database string
			Schema   string
		}{
			Host:     os.Getenv("DB_HOST"),
			Port:     port,
			User:     os.Getenv("DB_USER"),
			Password: os.Getenv("DB_PASSWORD"),
			Database: os.Getenv("DB_NAME"),
			Schema:   os.Getenv("DB_SCHEMA"),
		},
	}
}
