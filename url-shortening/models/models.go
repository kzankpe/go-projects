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
	Url        string    `json:"url" gorm:"column:url"`
	Shortcode  string    `json:"shortCode" gorm:"column:short_code;uniqueIndex"`
	CreatedAt  time.Time `json:"created_at"`
	UpdateddAt time.Time `json:"updated_at"`
	Count      int       `json:"count"`
}

type LongUrl struct {
	OUrl string `json:"url" validate:"required,notBlank"`
}

type Response struct {
	Data UrlData `json:"data"`
}

var DB *gorm.DB

func ConnectDB(conf config.Config) (*gorm.DB, error) {

	connectionStr := fmt.Sprintf("host=%s user=%s dbname=%s password=%s port=%s", conf.DBHost, conf.DBUserName, conf.DBName, conf.DBUserPass, conf.DBPort)

	db, err := gorm.Open(postgres.Open(connectionStr), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("error connecting to the database %w", err)
	}
	DB = db
	return db, nil
}
