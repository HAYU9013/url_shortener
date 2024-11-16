package repositories

import (
	"url_shortener/config"
	"url_shortener/models"
)

func CreateURL(url *models.URL) error {
	return config.DB.Create(url).Error
}

// GetURLByShortCode 根據短碼查詢對應的 URL 資料
// shortCode 是短碼的字串
// 返回值是 models.URL 結構和錯誤信息
func GetURLByShortCode(shortCode string) (models.URL, error) {
	var url models.URL
	err := config.DB.Where("short_code = ?", shortCode).First(&url).Error
	return url, err
}
