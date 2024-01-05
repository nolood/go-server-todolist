package postgres

import "github.com/google/uuid"

type Bill struct {
	ID      int       `json:"id"`
	UserID  uuid.UUID `pg:"type:uuid,alias:user_id" json:"user_id"`
	User    *User     `pg:"rel:has-one" json:"user"`
	Balance float32   `json:"balance"`
	Records []*Record `pg:"rel:has-many" json:"records"`
	Title   string    `pg:"unique:true" json:"title"`
}
