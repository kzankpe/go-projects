package models

import "time"

type Workout struct {
	ID           int
	UserId       int
	Name         string
	Description  string
	ScheduledFor time.Time
	//Exercises []Exercises
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
