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
	"gorm.io/gorm/clause"
)

type UrlController struct {
	DB *gorm.DB
}

func NewUrlController(DB *gorm.DB) UrlController {
	return UrlController{DB}
}

const req = "short_code = ?"

// Create short Url
// CreateShortUrl	godoc
//
//	@Summary		Create new short url
//	@Description	Create new short url
//	@Accept			json
//	@Produce		json
//	@Tags			Shortening URL
//	@Param			url	body	models.LongUrl	true	"Long Url"
//
//	@Sucess			200 {object} models.Response
//
//	@Router			/shorten [post]
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

	result := uc.DB.Create(&urldata).Clauses(clause.Returning{})
	if result.Error != nil {
		c.JSON(http.StatusBadGateway, gin.H{"status": "error", "message": result.Error.Error()})
		return
	}
	// response := models.Response{
	// 	ID:         urldata.ID,
	// 	Url:        urldata.Url,
	// 	Shortcode:  urldata.Shortcode,
	// 	CreatedAt:  urldata.CreatedAt,
	// 	UpdateddAt: urldata.UpdateddAt,
	// }
	//c.JSON(http.StatusCreated, gin.H{"id": response.ID, "url": response.Url, "shortcode": response.Shortcode, "created_at": response.CreatedAt, "updated_at": response.UpdateddAt})
	c.JSON(http.StatusCreated, models.Response{Data: urldata})
}

// Get short Url
//
//	@Summary		Retrieve  short url
//	@Description	Retrieve  short url
//	@Accept			json
//	@Produce		json
//	@Tags			Shortening URL
//	@Param			shortcode	path	string	true	"Short Code"
//
//	@Sucess			200 {object} models.Response
//
//	@Router			/shorten/{shortcode} [get]
func (uc *UrlController) RetrieveShortenUrl(c *gin.Context) {
	short := c.Param("shortcode")

	var url models.UrlData
	result := uc.DB.First(&url, req, short)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"status": "fail", "message": "No record with that Shortcode exists"})
		return
	}
	// Increment the count
	url.Count++
	result = uc.DB.Save(&url)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": "fail", "message": "Internal server Error"})
	}
	c.JSON(http.StatusOK, gin.H{"status": "success", "data": url})
}

// Update ShortUrl
//
//	@Summary		Update  short url
//	@Description	Update  short url
//	@Accept			json
//	@Produce		json
//	@Tags			Shortening URL
//	@Param			shortcode		body	models.UrlData	true	"Short Code"
//
//	@Sucess			200 {object} 	models.Response{}
//
//	@Router			/shorten/{shortcode} [put]
func (uc *UrlController) UpdateShortenUrl(c *gin.Context) {
	short := c.Param("shortcode")
	var payload *models.LongUrl
	err := c.ShouldBindJSON(&payload)
	if err != nil {
		c.JSON(http.StatusBadGateway, gin.H{"status": "fail", "message": err.Error()})
		return
	}

	var updatedUrl models.UrlData
	result := uc.DB.First(&updatedUrl, req, short)
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

//	@Summary		Delete  short url
//	@Description	Delete  short url
//	@Accept			json
//	@Produce		json
//	@Tags			Shortening URL
//	@Param			shortcode	path	string	true	"Short Code"
//
//	@Sucess			200 {object} models.Response{}
//
//	@Router			/shorten/{shortcode} [delete]
func (uc *UrlController) DeleteShortenUrl(c *gin.Context) {
	short := c.Param("shortcode")

	result := uc.DB.Delete(&models.UrlData{}, req, short)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"status": "fail", "message": "No record with that shortcode does not exists"})
	}
	c.JSON(http.StatusOK, gin.H{"status": "success"})
}

//	@Summary		Get   short url stats
//	@Description	Get  short url statistics
//	@Accept			json
//	@Produce		json
//	@Tags			Shortening URL
//	@Param			url	path	string	true	"Short Code"
//
//	@Sucess			200 {object} models.Response{}
//
//	@Router			/shorten/{shortcode}/stats [get]
func (uc *UrlController) GetShortenUrlStat(c *gin.Context) {
	short := c.Param("shortcode")
	var url models.UrlData
	result := uc.DB.First(&url, req, short)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"status": "fail", "message": "No record with that shortcode does not exists"})
	}
	c.JSON(http.StatusOK, gin.H{"status": "success", "data": url})

}
