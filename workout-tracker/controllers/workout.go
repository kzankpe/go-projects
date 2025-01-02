package controllers

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type WorkoutController struct {
	DB *gorm.DB
}

func NewWorkoutController(DB *gorm.DB) WorkoutController {
	return WorkoutController{DB}
}

func (wc *WorkoutController) CreateWorkout(c *gin.Context) {
	//function to create workoute
}

func (wc *WorkoutController) UpdateWorkout(c *gin.Context) {
	//update workout
}

func (wc *WorkoutController) DeleteWorkout(c *gin.Context) {
	//delete workout
}
