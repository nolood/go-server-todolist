package postgres

type Record struct {
	ID          int `json:"id"`
	ArticleID   int
	Article     Article `pg:"rel:has-one" json:"article"`
	Description string  `json:"description"`
	BillID      int     `json:"bill_id"`
	Bill        Bill    `pg:"rel:has-one" json:"bill"`
	Amount      int     `json:"amount"`
}
