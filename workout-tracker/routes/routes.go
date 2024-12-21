package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/kzankpe/go-projects/workout-tracker/controllers"
)

type AuthRouteController struct {
	authController controllers.AuthController
}

func NewAuthRouteController(authController controllers.AuthController) AuthRouteController {
	return AuthRouteController{authController}
}

func (arc *AuthRouteController) AuthRoute(rg *gin.RouterGroup) {
	router := rg.Group("/auth")

	router.POST("/register", arc.authController.SignUpUser)
	router.POST("/login", arc.authController.SignInUser)
}
