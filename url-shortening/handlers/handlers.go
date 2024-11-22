package handlers

import (
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/kzankpe/go-projects/url-shortening/helpers"
	"github.com/kzankpe/go-projects/url-shortening/models"
)

func CreateShortenUrl(c *gin.Context) {
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
		Shortcode:  helpers.GenerateShortCode(),
		CreatedAt:  time.Now().UTC(),
		UpdateddAt: time.Now().UTC(),
	}
	fmt.Println(urldata) // To be remove
	// Insert Information in the database

	c.JSON(http.StatusCreated, gin.H{"Message": "Url Created"})

	//return "Success"
}

func RetrieveShortenUrl(c *gin.Context) {
	//Retrieve information from url
	short := c.Param("shortcode")

	// sqlDB, err := db.DB()
	// if err != nil {
	// 	log.Fatal("Failed to get DB instance:", err)
	// }

	// err = sqlDB.Ping()
	// if err != nil {
	// 	log.Fatal("Database connection is not alive:", err)
	// }
	// var urls models.UrlData
	// err = db.Where().First(&urls, short).Error
	// if err != nil {
	// 	log.Fatal("Error fetching the url:", err)
	// }
}
