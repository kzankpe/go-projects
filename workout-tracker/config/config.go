package config

import (
	"fmt"

	"gorm.io/driver/sqlite"

	//"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Config struct {
	DBHost     string
	DBUserName string
	DBUserPass string
	DBName     string
	DBPort     string
}

var DB *gorm.DB

func ConnectDB(conf Config) (*gorm.DB, error) {

	//connectionStr := fmt.Sprintf("host=%s user=%s dbname=%s password=%s port=%s", conf.DBHost, conf.DBUserName, conf.DBName, conf.DBUserPass, conf.DBPort)

	//db, err := gorm.Open(postgres.Open(connectionStr), &gorm.Config{})
	db, err := gorm.Open(sqlite.Open("gorm.db"), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("error connecting to the database %w", err)
	}
	DB = db
	return db, nil
}
