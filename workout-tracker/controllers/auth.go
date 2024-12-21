package controllers

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type AuthController struct {
	DB *gorm.DB
}

func NewAuthController(DB *gorm.DB) AuthController {
	return AuthController{DB}
}

//SignUp user

func (ac *AuthController) SignUpUser(c *gin.Context) {
	//signup logic
}

//SignIn User

func (ac *AuthController) SignInUser(c *gin.Context) {
	//SignIn logic
}
