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
	router.GET("/shorten/:shortcode", urc.urlController.RetrieveShortenUrl)
	router.PUT("/shorten/:shortcode", urc.urlController.UpdateShortenUrl)
	router.DELETE("/shorten/:shortcode", urc.urlController.DeleteShortenUrl)
}

