package main

import (
	"url_shortener/config"
	"url_shortener/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	// 答案在這個 branch
	// 初始化資料庫
	config.ConnectDatabase()
	// 初始化 Gin
	r := gin.Default()
	// 設定路由
	routes.SetupRoutes(r)
	// 啟動服務
	r.Run(":8080")
}
