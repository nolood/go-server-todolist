package postgres

import (
	"gorm.io/gorm"
)

type Bill struct {
	gorm.Model
	ID      uint64 `gorm:"autoIncrement"`
	User    User   `gorm:"foreignKey:UserID"`
	UserID  uint64
	Balance float32
	Records []Record
	Title   string
}
