package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/kzankpe/go-projects/url-shortening/handlers"
)

type UrlRouteController struct {
	urlController handlers.UrlController
}

func NewRouteUrlController(uc handlers.UrlController) UrlRouteController {
	return UrlRouteController{uc}
}
func (urc *UrlRouteController) UrlRoute(rg *gin.RouterGroup) {
	router := rg.Group("/v1")
	//router.Use(middleware.DeserializeUser())
	router.POST("/shorten", urc.urlController.CreateShortenUrl)
}

func InitRouter() *gin.Engine {
	router := gin.Default()
	return router
}

func SetRoute(route *gin.Engine) {
	// Create a group based on the endpiont version
	v1 := route.Group("/v1")
	{
		v1.GET("/shorten/:shortCode", handlers.RetrieveShortenUrl)
	}

}
