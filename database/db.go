package database

import (
	"github.com/similadayo/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func init() {
	// initialize the databas
	db, err := gorm.Open(sqlite.Open("database/Liveshare.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// Migrate the schema
	db.AutoMigrate(&models.User{})
}
