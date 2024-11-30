package handlers

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/kzankpe/go-projects/url-shortening/helpers"
	"github.com/kzankpe/go-projects/url-shortening/models"
	"gorm.io/gorm"
)

type UrlController struct {
	DB *gorm.DB
}

func NewUrlController(DB *gorm.DB) UrlController {
	return UrlController{DB}
}

// Create short Url
func (uc *UrlController) CreateShortenUrl(c *gin.Context) {
	var newUrl models.OriginalUrl
	err := c.ShouldBindJSON(&newUrl)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	// Create url info
	urldata := models.UrlData{
		Url:        strings.ToLower(strings.TrimSpace(newUrl.OUrl)),
		Shortcode:  helpers.GenerateShortCode(),
		CreatedAt:  time.Now().UTC(),
		UpdateddAt: time.Now().UTC(),
	}

	result := uc.DB.Create(&urldata)
	if result.Error != nil {
		c.JSON(http.StatusBadGateway, gin.H{"status": "error", "message": result.Error.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"status": "success", "data": urldata})
}

func RetrieveShortenUrl(c *gin.Context) {
	//Retrieve information from url
	short := c.Param("shortcode")
	fmt.Println(short)
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
