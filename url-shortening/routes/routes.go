package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/kzankpe/go-projects/url-shortening/handlers"
)

func InitRouter() *gin.Engine {
	router := gin.Default()
	return router
}

func SetRoute(route *gin.Engine) {
	// Create a group based on the endpiont version
	v1 := route.Group("/v1")
	{
		v1.POST("/shorten", handlers.CreateShortenUrl)
		v1.GET("/shorten", handlers.RetrieveShortenUrl)
	}

}
