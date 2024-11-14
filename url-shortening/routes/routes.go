package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/kzankpe/go-projects/url-shortening/handlers"
)

func SetUpRoute(route *gin.Engine) {
	route.POST("/shorten", handlers.ShortenUrl)
}
