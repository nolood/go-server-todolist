package postgres

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID         uint64 `gorm:"autoIncrement"`
	Username   string
	Password   string
	TelegramID int64 `json:"telegram_id"`
	VKID       int64 `json:"vk_id"`
	Email      string
	Phone      string
	Bills      []Bill
}
