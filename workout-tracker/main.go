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
	_ "github.com/kzankpe/go-projects/workout-tracker/docs"
	"github.com/kzankpe/go-projects/workout-tracker/models"
	"github.com/kzankpe/go-projects/workout-tracker/routes"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

var (
	server                  *gin.Engine
	AuthController          controllers.AuthController
	AuthRouteController     routes.AuthRouteController
	ExerciseController      controllers.ExerciseController
	ExerciseRouteController routes.ExerciseRouteController
	WorkoutController       controllers.WorkoutController
	WorkoutRouteController  routes.WorkoutRouteController
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
	//Initializing server
	server = gin.Default()
	fmt.Println("Successfully initialize the server ....")
}

// @title			Workout Tracker Service
// @version		1.0
// @description	A simple Service using RESTful API that allows users to track their workouts
// @version		1.0
// @contact.name	kzankpe
// @contact.url	https://github.com/kzankpe
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
// @BasePath		/api/v1
func main() {
	//main function
	router := server.Group("/api")
	router.GET("/healthcheck", func(c *gin.Context) {
		message := "Welcome to Url Shortening service"
		c.JSON(http.StatusOK, gin.H{"status": "success", "message": message})
	})
	AuthRouteController.AuthRoute(router)
	ExerciseRouteController.ExerciseRoute(router)
	WorkoutRouteController.WorkoutRoute(router)
	// Add Swagger
	router.GET("/docs/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	// Run the server
	err := server.Run(":8090")
	if err != nil {
		panic(err)
	}
}
