package middleware

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/kzankpe/go-projects/workout-tracker/config"
	"github.com/kzankpe/go-projects/workout-tracker/helper"
	"github.com/kzankpe/go-projects/workout-tracker/models"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		var accessToken string
		authHeader := c.Request.Header.Get("Authorization")

		if authHeader == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"status": "fail", "message": "Missing authorization header"})
			return
		}

		fields := strings.Fields(authHeader)
		if len(fields) != 0 && fields[0] == "Bearer" {
			accessToken = fields[1]
		} else {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"status": "fail", "message": "Missing authorization header"})
			return
		}
		userId, err := helper.ValidateJWT(accessToken, "") // TODO add privatekey here
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"status": "fail", "message": "Invalid Token"})
			return
		}
		fmt.Println(userId)
		// Verify the userId in the database
		var user models.User
		result := config.DB.First(&user, "id= ? ", fmt.Sprint(userId))
		if result.Error !=nil {
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"status": "fail", "message": "the user belonging to this token do not exist"})
			return
		}
		c.Set("currentUser", userId)
		c.Next()
	}
}
