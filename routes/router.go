package routes

import (
	"url_shortener/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {
	r.Static("/public", "./public")
	r.POST("/shorten", controllers.ShortenURL)
	r.GET("/r/:code", controllers.RedirectURL)
}
