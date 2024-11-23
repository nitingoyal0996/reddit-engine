package database

import (
	"log"

	"github.com/nitingoyal0996/reddit-clone/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func InitDB() (*gorm.DB, error) {
	dsn := "file:reddit-engine.db?_pragma=key('engine')"
	// add sqlite database with passphrase
	db, err := gorm.Open(sqlite.Open(dsn))
	if err != nil {
		log.Fatalf("failed to open database: %v", err)
		return nil, err
	}
	// auto migrate
	err = db.AutoMigrate(&models.User{}, &models.Message{})
	return db, err
}
