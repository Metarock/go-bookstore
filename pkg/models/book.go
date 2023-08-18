package models

import (
	"github.com/jinzhu/gorm"

	"github.com/Metarock/go-bookstore/pkg/config"
)

var db *gorm.DB

type Book struct {
	gorm.Model
	Name        string `json:"name"`
	Author      string `json:"author"`
	Publication string `json:"publication"`
}

func init() {
	// from the conig/app.go
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&Book{})
}
