package postgres

import "gorm.io/gorm"

type ArticleType struct {
	ID   uint64 `gorm:"autoIncrement"`
	Name string
}

type Article struct {
	gorm.Model
	ID            uint64 `gorm:"autoIncrement"`
	Icon          string
	ArticleTypeID uint
	ArticleType   ArticleType
}
