package postgres

import "gorm.io/gorm"

type Record struct {
	gorm.Model
	ID          uint64 `gorm:"autoIncrement"`
	ArticleID   int
	Article     Article
	Description string
	BillID      int
	Bill        Bill
	Amount      int
}
