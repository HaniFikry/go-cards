package db

import (
	"github.com/HaniFikry/go-cards/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Init() {
	db, err := gorm.Open(sqlite.Open("cards.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(&models.Deck{})
	DB = db
}

func GetDB() *gorm.DB {
	return DB
}
