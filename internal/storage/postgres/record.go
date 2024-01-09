package postgres

import (
	"gorm.io/gorm"
	"time"
)

type RecordType struct {
	gorm.Model
	ID    uint64 `gorm:"autoIncrement" json:"id"`
	Value string `json:"value"`
}

type Record struct {
	gorm.Model
	ID          uint64     `gorm:"autoIncrement" json:"id"`
	Article     Article    `json:"article"`
	ArticleID   uint64     `json:"article_id"`
	Description string     `json:"description"`
	Bill        Bill       `json:"bill"`
	BillID      uint64     `json:"bill_id"`
	Amount      int        `json:"amount"`
	Type        RecordType `json:"type"`
	TypeId      uint64     `json:"type_id"`
	Date        time.Time  `json:"date"`
}
