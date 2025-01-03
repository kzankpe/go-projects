package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/kzankpe/go-projects/workout-tracker/controllers"
	"github.com/kzankpe/go-projects/workout-tracker/middleware"
)

type WorkoutRouteController struct {
	WorkoutController controllers.WorkoutController
}

func NewWorkoutRouteController(workoutController controllers.WorkoutController) WorkoutRouteController {
	return WorkoutRouteController{workoutController}
}

func (wrc *WorkoutRouteController) WorkoutRoute(rg *gin.RouterGroup) {
	router := rg.Group("/workout", middleware.AuthMiddleware())
	router.POST("/create", wrc.WorkoutController.CreateWorkout)
	router.PUT("/:workoutid", wrc.WorkoutController.UpdateWorkout)
	router.DELETE("/:workoutid", wrc.WorkoutController.DeleteWorkout)
}
