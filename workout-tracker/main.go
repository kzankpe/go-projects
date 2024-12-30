package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	"github.com/kzankpe/go-projects/workout-tracker/config"
	"github.com/kzankpe/go-projects/workout-tracker/controllers"
	"github.com/kzankpe/go-projects/workout-tracker/models"
	"github.com/kzankpe/go-projects/workout-tracker/routes"
)

var (
	server                  *gin.Engine
	AuthController          controllers.AuthController
	AuthRouteController     routes.AuthRouteController
	ExerciseController      controllers.ExerciseController
	ExerciseRouteController routes.ExerciseRouteController
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Could not load environment variables: %v", err)
	}

	// Database configuration
	conf := config.Config{
		DBHost:     os.Getenv("DB_HOST"),
		DBUserName: os.Getenv("DB_USER"),
		DBUserPass: os.Getenv("DB_PASSWORD"),
		DBName:     os.Getenv("DB_NAME"),
		DBPort:     os.Getenv("DB_PORT"),
	}
	db, err := config.ConnectDB(conf)
	if err != nil {
		log.Fatalf("Failed to connect to the databse %v", err)
	}
	//Migrate database
	err = db.AutoMigrate(&models.User{})
	if err != nil {
		log.Fatalf("Error migrating the database: %v", err)
	}
	fmt.Println("Successfully connected to the postgresql database!")

}

func main() {
	//main function
	router := server.Group("/api")
	router.GET("/healthcheck", func(c *gin.Context) {
		message := "Welcome to Url Shortening service"
		c.JSON(http.StatusOK, gin.H{"status": "success", "message": message})
	})
	AuthRouteController.AuthRoute(router)
	ExerciseRouteController.ExerciseRoute(router)
	// Run the server
	err := server.Run(":8090")
	if err != nil {
		panic(err)
	}
}
