package main

import "github.com/gin-gonic/gin"

func main() {
	// Create a new router
	router := gin.Default()

	// Define a handler route for root
	router.GET("/", func(ctx *gin.Context) {
		// Send a Json response
		ctx.JSON(200, gin.H{
			"mesage": "Hello from Gin!",
		})
	})

	// Run the server on a port
	err := router.Run(":8090")
	if err != nil {
		panic(err)
	}
}
