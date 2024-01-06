package postgres

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID         uint64 `gorm:"autoIncrement" json:"id"`
	Username   string `json:"username"`
	Password   string `json:"password"`
	TelegramID int64  `json:"telegram_id"`
	VKID       int64  `json:"vk_id"`
	Email      string `json:"email"`
	Phone      string `json:"phone"`
	Bills      []Bill `json:"bills"`
}
