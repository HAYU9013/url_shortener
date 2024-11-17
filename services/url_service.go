package services

import (
	//TODO: import math/rand package
	"url_shortener/models"       // 引入自定義的models包
	"url_shortener/repositories" // 引入自定義的repositories包
)

func GenerateShortCode() string {
	chars := "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789-_" // 定義短碼的字元集
	var b string
	// TODO: Generate a 6-character short code save into b

	return b
}

func ShortenURL(originalURL string) (models.URL, error) {
	shortCode := GenerateShortCode()                                  // 生成短碼
	url := models.URL{OriginalURL: originalURL, ShortCode: shortCode} // 創建URL模型實例
	err := repositories.CreateURL(&url)                               // 將URL模型實例保存到數據庫
	return url, err                                                   // 返回URL模型實例和錯誤信息
}
