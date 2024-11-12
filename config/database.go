package config

import (
	"log"
	"url_shortener/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	dsn := "root:PPAASSWW@tcp(127.0.0.1:3306)/url_shortener?charset=utf8mb4&parseTime=True&loc=Local"
	database, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}
	DB = database
	DB.AutoMigrate(&models.URL{})
}
