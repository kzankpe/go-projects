package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func ShortenUrl(c *gin.Context) {

	c.JSON(http.StatusCreated, gin.H{"Message": "Url Created"})
}
