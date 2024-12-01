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
	var newUrl models.LongUrl
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

// Get short Url
func (uc *UrlController) RetrieveShortenUrl(c *gin.Context) {
	short := c.Param("shortcode")

	var url models.UrlData
	result := uc.DB.First(&url, "shortcode = ?", short)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"status": "fail", "message": "No record with that Shortcode exists"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "success", "data": url})
}

// Update ShortUrl
func (uc *UrlController) UpdateShortenUrl(c *gin.Context) {
	short := c.Param("shortcode")
	var payload *models.LongUrl
	err := c.ShouldBindJSON(&payload)
	if err != nil {
		c.JSON(http.StatusBadGateway, gin.H{"status": "fail", "message": err.Error()})
		return
	}

	var updatedUrl models.UrlData
	result := uc.DB.First(&updatedUrl, "shortCode = ?", short)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"status": "fail", "message": "No record with that Shortcode exists"})
		return
	}
	updatedInput := models.UrlData{
		Url:        payload.OUrl,
		Shortcode:  updatedUrl.Shortcode,
		CreatedAt:  updatedUrl.CreatedAt,
		UpdateddAt: time.Now().UTC(),
	}

	uc.DB.Model(&updatedUrl).Updates(updatedInput)

	c.JSON(http.StatusOK, gin.H{"status": "success", "data": updatedUrl})
}

func (uc *UrlController) DeleteShortenUrl(c *gin.Context) {
	short := c.Param("shortcode")

	result := uc.DB.Delete(&models.UrlData{}, "shortCode = ?", short)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"status": "fail", "message": "No record with that shortcode exists"})
	}
}
