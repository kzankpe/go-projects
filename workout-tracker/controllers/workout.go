package controllers

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/kzankpe/go-projects/workout-tracker/models"
	"gorm.io/gorm"
)

type WorkoutController struct {
	DB *gorm.DB
}

func NewWorkoutController(DB *gorm.DB) WorkoutController {
	return WorkoutController{DB}
}

func (wc *WorkoutController) CreateWorkout(c *gin.Context) {
	//function to create workout

	var payload *models.Workout
	err := c.ShouldBindJSON(&payload)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": err.Error()})
		return
	}

	newWorkout := models.Workout{
		Name:         payload.Name,
		Description:  payload.Description,
		ScheduledFor: payload.ScheduledFor,
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	}

	result := wc.DB.Create(&newWorkout)
	if result.Error != nil {
		c.JSON(http.StatusBadGateway, gin.H{"status": "error", "message": "Something bad happened"})
		return
	}
}

func (wc *WorkoutController) UpdateWorkout(c *gin.Context) {
	//update workout
}

func (wc *WorkoutController) DeleteWorkout(c *gin.Context) {
	//delete workout
}
