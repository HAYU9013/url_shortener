package services

import (
	"crypto/rand"
	"encoding/base64"
	"url_shortener/models"
	"url_shortener/repositories"
)

func GenerateShortCode() string {
	b := make([]byte, 6)
	rand.Read(b)
	return base64.URLEncoding.EncodeToString(b)
}

func ShortenURL(originalURL string) (models.URL, error) {
	shortCode := GenerateShortCode()
	url := models.URL{OriginalURL: originalURL, ShortCode: shortCode}
	err := repositories.CreateURL(&url)
	return url, err
}
