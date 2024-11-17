package routes

import (
	// TODO: import url_shortener/controllers package
	"github.com/gin-gonic/gin"
)

// SetupRoutes 設置應用程式的路由
// r *gin.Engine 是 Gin 框架的引擎實例
func SetupRoutes(r *gin.Engine) {
	r.Static("/public", "./public")
	// TODO: set POST method of /shorten for shorten url, call controllers.ShortenURL

	// TODO: set GET method of /r/:code for redirect url, call controllers.RedirectURL

}
