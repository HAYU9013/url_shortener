package config

import (
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

// ConnectDatabase 連接到 MySQL 資料庫，並進行資料庫遷移
// 使用 GORM 來打開資料庫連接
// 如果連接失敗，會記錄錯誤並終止程序

func ConnectDatabase() {
	dsn := "root:PPAASSWW@tcp(127.0.0.1:3306)/url_shortener?charset=utf8mb4&parseTime=True&loc=Local"
	database, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}
	DB = database
}
