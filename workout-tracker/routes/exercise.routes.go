package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/kzankpe/go-projects/workout-tracker/controllers"
)

type ExerciseRouteController struct {
	ExerciseController controllers.ExerciseController
}

func NewExerciseRouteController(exerciseController controllers.ExerciseController) ExerciseRouteController {
	return ExerciseRouteController{exerciseController}
}

func (erc *ExerciseRouteController) ExerciseRoute(rg *gin.RouterGroup) {
	router := rg.Group("/exercise")

	router.POST("/create", erc.ExerciseController.CreateExercise)
	router.PUT("/:exercise", erc.ExerciseController.UpdateExercise)
	router.DELETE("/:exercise", erc.ExerciseController.DeleteExercise)
}
