package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/kzankpe/go-projects/url-shortening/config"
	"github.com/kzankpe/go-projects/url-shortening/handlers"
	"github.com/kzankpe/go-projects/url-shortening/models"
	"github.com/kzankpe/go-projects/url-shortening/routes"
)

var (
	server             *gin.Engine
	UrlController      handlers.UrlController
	UrlRouteController routes.UrlRouteController
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Could not load environment variables: %v", err)
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

	db, err := models.ConnectDB(config)
	if err != nil {
		log.Fatalf("Failed to connect to the databse %v", err)
	}

	//Migrate database
	err = db.AutoMigrate(&models.UrlData{})
	if err != nil {
		log.Fatalf("Error migrating the database: %v", err)
	}
	fmt.Println("Successfully connected to the database!")

	UrlController = handlers.NewUrlController(db)
	UrlRouteController = routes.NewRouteUrlController(UrlController)
	//Initializing server
	server = gin.Default()

}

func main() {
	router := server.Group("/api")
	router.GET("/healthcheck", func(c *gin.Context) {
		message := "Welcome to Url Shortening service"
		c.JSON(http.StatusOK, gin.H{"status": "success", "message": message})
	})
	//Add routes
	UrlRouteController.UrlRoute(router)
	// Run the server on a port
	err := server.Run(":8090")
	if err != nil {
		panic(err)
	}
}
