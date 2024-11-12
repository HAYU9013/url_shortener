package repositories

import (
	"url_shortener/config"
	"url_shortener/models"
)

func CreateURL(url *models.URL) error {
	return config.DB.Create(url).Error
}

func GetURLByShortCode(shortCode string) (models.URL, error) {
	var url models.URL
	err := config.DB.Where("short_code = ?", shortCode).First(&url).Error
	return url, err
}
