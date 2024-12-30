package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kzankpe/go-projects/workout-tracker/models"
	"gorm.io/gorm"
)

type ExerciseController struct {
	DB *gorm.DB
}

func NewExerciseController(DB *gorm.DB) ExerciseController {
	return ExerciseController{DB}
}

func (ec *ExerciseController) CreateExercise(c *gin.Context) {
	var payload *models.UserExerciseInput

	err := c.ShouldBindJSON(&payload)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": err.Error()})
		return
	}

	newExercise := models.Exercise{
		Name:        payload.Name,
		Description: payload.Description,
		Category:    payload.Category,
		MuscleGroup: payload.MuscleGroup,
	}

	result := ec.DB.Create(&newExercise)
	if result.Error != nil {
		c.JSON(http.StatusBadGateway, gin.H{"status": "error", "message": "Something bad happened"})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"status": "success", "data": gin.H{"exercise": "created"}})
}

func (ec *ExerciseController) UpdateExercise(c *gin.Context) {
	// update an existing exercise in the database
}

func (ec *ExerciseController) DeleteExercise(c *gin.Context) {
	// Delete exercise from database
}
