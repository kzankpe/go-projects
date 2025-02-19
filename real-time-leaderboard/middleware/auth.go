package middleware

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		var accessToken string
		authHeader := c.Request.Header.Get("Authorization")

		// Check if the authHeader is empty and return an error
		if authHeader == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"status": "fail", "message": "Missing authorization header"})
			return
		}

		fields := strings.Fields(authHeader)
		if len(fields) != 0 && fields[0] == "Bearer" {
			accessToken = fields[1]
		} else {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"status": "failed", "message": "Missing authorization header"})
		}
		// Verify the Token
		fmt.Println(accessToken) // To discard var unused error
		// Verifiy the user in the database

	}
}
