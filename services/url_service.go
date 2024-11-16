package services

import (
	"math/rand"                  // 引入math/rand包
	"url_shortener/models"       // 引入自定義的models包
	"url_shortener/repositories" // 引入自定義的repositories包
)

func GenerateShortCode() string {
	chars := "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789-_" // 定義短碼的字符集
	var b string
	for i := 0; i < 6; i++ {
		b += string(chars[rand.Intn(len(chars))]) // 創建一個長度為6的字節數組
	}
	return b
}

func ShortenURL(originalURL string) (models.URL, error) {
	shortCode := GenerateShortCode()                                  // 生成短碼
	url := models.URL{OriginalURL: originalURL, ShortCode: shortCode} // 創建URL模型實例
	err := repositories.CreateURL(&url)                               // 將URL模型實例保存到數據庫
	return url, err                                                   // 返回URL模型實例和錯誤信息
}
