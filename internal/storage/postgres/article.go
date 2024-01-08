package postgres

import "gorm.io/gorm"

type Article struct {
	gorm.Model
	ID     uint64 `gorm:"autoIncrement"`
	Icon   string
	Title  string
	Color  string
	UserID uint64
	User   User
}
