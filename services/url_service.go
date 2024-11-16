package services

import (
	"crypto/rand"                // 引入隨機數生成包
	"encoding/base64"            // 引入base64編碼包
	"url_shortener/models"       // 引入自定義的models包
	"url_shortener/repositories" // 引入自定義的repositories包
)

func GenerateShortCode() string {
	b := make([]byte, 6)                        // 創建一個6字節的切片
	rand.Read(b)                                // 生成隨機數並填充切片
	return base64.URLEncoding.EncodeToString(b) // 將隨機數編碼為base64並返回
}

func ShortenURL(originalURL string) (models.URL, error) {
	shortCode := GenerateShortCode()                                  // 生成短碼
	url := models.URL{OriginalURL: originalURL, ShortCode: shortCode} // 創建URL模型實例
	err := repositories.CreateURL(&url)                               // 將URL模型實例保存到數據庫
	return url, err                                                   // 返回URL模型實例和錯誤信息
}
