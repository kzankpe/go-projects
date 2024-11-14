package models

import "time"

type UrlData struct {
	ID         int       `json:"id"`
	Url        string    `json:"url"`
	Shortcode  string    `json:"shortCode"`
	CreatedAt  time.Time `json:"created_at"`
	UpdateddAt time.Time `json:"updated_at"`
	Count      int       `json:"count"`
}

type OriginalUrl struct {
	OUrl string `json:"url" validate:"required,notBlank"`
}
