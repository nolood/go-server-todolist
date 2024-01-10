package postgres

type Bill struct {
	Model
	User    User     `gorm:"foreignKey:UserID" json:"user"`
	UserID  uint64   `json:"user_id"`
	Balance float32  `json:"balance"`
	Records []Record `json:"records"`
	Title   string   `json:"title"`
}
