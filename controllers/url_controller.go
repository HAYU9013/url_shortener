package controllers

import (
	"fmt"
	"net/http"
	"url_shortener/repositories"
	"url_shortener/services"

	"github.com/gin-gonic/gin"
)

// 縮短URL的處理函數
func ShortenURL(c *gin.Context) {
	// 從請求中獲取原始URL
	var requestBody struct {
		URL string `json:"url"`
	}
	// TODO: Bind request body (json) to requestBody

	originalURL := requestBody.URL
	// 打印原始URL
	fmt.Println(originalURL)
	// 調用服務來縮短URL
	url, err := services.ShortenURL(originalURL)
	// 如果出錯，返回內部服務器錯誤
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to shorten URL"})
		return
	}
	// 設定端口號
	port := "8080"
	// TODO: Send response with shortened URL

}

// 重定向URL的處理函數
func RedirectURL(c *gin.Context) {

	// TODO: Get short code from request URL route parameter

	// 打印短碼
	fmt.Println(shortCode)
	// 從存儲庫中獲取原始URL
	url, err := repositories.GetURLByShortCode(shortCode)
	// 如果出錯，返回未找到錯誤
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "URL not found"})
		return
	}
	// 重定向到原始URL
	c.Redirect(http.StatusMovedPermanently, url.OriginalURL)
}
