package models

import "gorm.io/gorm"

// URL 結構體代表一個短網址的模型
// 包含原始網址和短碼
type URL struct {
	// gorm.Model 提供基本的模型字段，如 ID、創建時間和更新時間
	gorm.Model
	// OriginalURL 是原始的長網址，不能為空
	OriginalURL string `gorm:"not null"`
	// ShortCode 是短網址的唯一代碼，必須唯一且不能為空
	ShortCode string `gorm:"uniqueIndex;not null"`
}
