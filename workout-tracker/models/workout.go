package models

import (
	"time"

	"gorm.io/gorm"
)

type Workout struct {
	gorm.Model
	UserId       int
	Name         string
	Description  string
	ScheduledFor time.Time
	//Exercises    []Exercise
}
