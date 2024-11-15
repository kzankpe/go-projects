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
	route.POST("/shorten", handlers.ShortenUrl)
}
