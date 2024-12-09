package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/kzankpe/go-projects/url-shortening/config"
	_ "github.com/kzankpe/go-projects/url-shortening/docs"
	"github.com/kzankpe/go-projects/url-shortening/handlers"
	"github.com/kzankpe/go-projects/url-shortening/models"
	"github.com/kzankpe/go-projects/url-shortening/routes"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
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

//	@title			Shorten Url Service
//	@version		1.0
//	@description	A simple RESTful API that allows users to shorten long URLs
//	@version		1.0
//	@contact.name	kzankpe
//	@contact.url	https://github.com/kzankpe
//	@BasePath		/api/v1
func main() {
	router := server.Group("/api")
	router.GET("/healthcheck", func(c *gin.Context) {
		message := "Welcome to Url Shortening service"
		c.JSON(http.StatusOK, gin.H{"status": "success", "message": message})
	})
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	//Add routes
	UrlRouteController.UrlRoute(router)
	// Run the server on a port
	err := server.Run(":8090")
	if err != nil {
		panic(err)
	}
}
