package models

import "github.com/google/uuid"

type CreateBillDto struct {
	tableName struct{}  `pg:"bills"`
	ID        int       `json:"id"`
	UserID    uuid.UUID `json:"user_id"`
	Balance   float32   `json:"balance"`
	Title     string    `json:"title"`
}
