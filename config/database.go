package config

import (
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	"GoBlogClean/models"
)

type DBHandler struct {
	Conn *gorm.DB
}

func NewDBHandler() *DBHandler {
	conn, err := gorm.Open(sqlite.Open("gorm.db"), &gorm.Config{})
	if err != nil {
		log.Println(err)
		return nil
	}

	if err = conn.AutoMigrate(&models.Article{}, &models.Comment{}, &models.User{}); err != nil {
		log.Println(err)
		return nil
	}

	dbHandler := new(DBHandler)
	dbHandler.Conn = conn
	return dbHandler
}
