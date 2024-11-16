package models

import (
	"fmt"
	"time"

	"github.com/kzankpe/go-projects/url-shortening/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type UrlData struct {
	ID         uint      `json:"id" gorm:"primaryKey"`
	Url        string    `json:"url"`
	Shortcode  string    `json:"shortCode"`
	CreatedAt  time.Time `json:"created_at"`
	UpdateddAt time.Time `json:"updated_at"`
	Count      int       `json:"count"`
}

type OriginalUrl struct {
	OUrl string `json:"url" validate:"required,notBlank"`
}

func ConnectDB(conf config.Config) (*gorm.DB, error) {
	connectionStr := fmt.Sprintf("host=%s user=%s dbname=%s password=%s port=%s sslmode=%s", conf.DBHost, conf.DBUserName, conf.DBName, conf.DBUserPass, conf.DBPort, conf.DBSslMode)

	db, err := gorm.Open(postgres.Open(connectionStr), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("error connecting to the database %w", err)
	}
	return db, nil
}
