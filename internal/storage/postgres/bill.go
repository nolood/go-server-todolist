package postgres

import (
	"gorm.io/gorm"
)

type Bill struct {
	gorm.Model
	ID      uint64   `gorm:"autoIncrement" json:"id"`
	User    User     `gorm:"foreignKey:UserID" json:"user"`
	UserID  uint64   `json:"user_id"`
	Balance float32  `json:"balance"`
	Records []Record `json:"records"`
	Title   string   `json:"title"`
}
