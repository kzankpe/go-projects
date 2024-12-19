package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/kzankpe/go-projects/workout-tracker/config"
	"github.com/kzankpe/go-projects/workout-tracker/models"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Could not load environment variables: %v", err)
	}

	// Database configuration
	config := config.Config{
		DBHost:     os.Getenv("DB_HOST"),
		DBUserName: os.Getenv("DB_USER"),
		DBUserPass: os.Getenv("DB_PASSWORD"),
		DBName:     os.Getenv("DB_NAME"),
		DBPort:     os.Getenv("DB_PORT"),
		DBSslMode:  os.Getenv("DB_SSLMODE"),
	}
	db, err := models.ConnectDB(config)
	if err != nil {
		log.Fatalf("Failed to connect to the databse %v", err)
	}
	//Migrate database
	err = db.AutoMigrate(&models.User{})
	if err != nil {
		log.Fatalf("Error migrating the database: %v", err)
	}
	fmt.Println("Successfully connected to the database!")

}

func main() {

}
