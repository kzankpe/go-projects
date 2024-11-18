package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/kzankpe/go-projects/url-shortening/config"
	"github.com/kzankpe/go-projects/url-shortening/models"
	"github.com/kzankpe/go-projects/url-shortening/routes"
)

func main() {
	// Load environment variables from .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}
	// Database configuration from env
	config := config.Config{
		DBHost:     os.Getenv("DB_HOST"),
		DBUserName: os.Getenv("DB_USER"),
		DBUserPass: os.Getenv("DB_PASSWORD"),
		DBName:     os.Getenv("DB_NAME"),
		DBPort:     os.Getenv("DB_PORT"),
		DBSslMode:  os.Getenv("DB_SSLMODE"),
	}

	// Initialize a router
	r := routes.InitRouter()
	routes.SetRoute(r)
	// Connect to the Database
	db, err := models.ConnectDB(config)
	if err != nil {
		log.Fatalf("Failed to connect to the databse %v", err)
	}
	err = db.AutoMigrate(&models.UrlData{})
	if err != nil {
		log.Fatalf("Error migrating the database: %v", err)
	}
	fmt.Println("Successfully connected to the database!")

	// Run the server on a port
	err = r.Run(":8090")
	if err != nil {
		panic(err)
	}
}
