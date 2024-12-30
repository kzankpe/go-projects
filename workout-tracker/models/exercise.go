package models

import "gorm.io/gorm"

type Exercise struct {
	gorm.Model
	Name        string `json:"name"`
	Description string `json:"description"`
	Category    string `json:"category"`
	MuscleGroup string `json:"muscle_group"`
}

type UserExerciseInput struct {
	Name        string
	Description string
	Category    string
	MuscleGroup string
}
