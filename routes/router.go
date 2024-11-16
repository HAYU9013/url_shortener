package routes

import (
	"url_shortener/controllers"

	"github.com/gin-gonic/gin"
)

// SetupRoutes 設置應用程式的路由
// r *gin.Engine 是 Gin 框架的引擎實例
// r.Static("/public", "./public") 設置靜態文件的路徑
// r.POST("/shorten", controllers.ShortenURL) 設置縮短 URL 的 POST 路由
// r.GET("/r/:code", controllers.RedirectURL) 設置重定向 URL 的 GET 路由
func SetupRoutes(r *gin.Engine) {
	r.Static("/public", "./public")
	r.POST("/shorten", controllers.ShortenURL)
	r.GET("/r/:code", controllers.RedirectURL)
}
