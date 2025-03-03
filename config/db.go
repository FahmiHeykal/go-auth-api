package config

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var DB *gorm.DB

func InitDB() (*gorm.DB, error) {
	db, err := gorm.Open("postgres", "host=localhost port=5432 user=postgres dbname=go_auth password=1234 sslmode=disable")
	if err != nil {
		return nil, err
	}
	DB = db
	return DB, nil
}
