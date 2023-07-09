package db

import (
	"web-service-gin/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func Init() *gorm.DB {
	dsn := "tutorial-go:udPHP_r68.LQ@tcp(127.0.0.1:3306)/tutorial_go?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}
	db.AutoMigrate(&models.Album{})
	return db
}
