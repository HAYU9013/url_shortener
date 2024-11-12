package controllers

import (
	"fmt"
	"net/http"
	"url_shortener/repositories"
	"url_shortener/services"

	"github.com/gin-gonic/gin"
)

func ShortenURL(c *gin.Context) {
	originalURL := c.PostForm("url")
	fmt.Println(originalURL)
	url, err := services.ShortenURL(originalURL)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to shorten URL"})
		return
	}
	port := "8080"
	c.JSON(http.StatusOK, gin.H{"short_url": "http://localhost:" + port + "/r/" + url.ShortCode})
}

func RedirectURL(c *gin.Context) {
	shortCode := c.Param("code")
	fmt.Println(shortCode)
	url, err := repositories.GetURLByShortCode(shortCode)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "URL not found"})
		return
	}
	c.Redirect(http.StatusMovedPermanently, url.OriginalURL)
}
