package handlers

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/kzankpe/go-projects/url-shortening/models"
)

func ShortenUrl(c *gin.Context) {
	var newUrl models.OriginalUrl
	// Retrieve the url information from the
	err := c.BindJSON(&newUrl)

	if err != nil {
		fmt.Println(err)
		return //http.ErrBodyNotAllowed
	}

	// Use the validor function here before continuing

	// Create url info
	urldata := models.UrlData{
		Url:        strings.ToLower(strings.TrimSpace(newUrl.OUrl)),
		Shortcode:  "",
		CreatedAt:  time.Now().UTC(),
		UpdateddAt: time.Now().UTC(),
	}
	fmt.Println(urldata) // To be remove
	// Insert Information in the database

	c.JSON(http.StatusCreated, gin.H{"Message": "Url Created"})

	//return "Success"
}
