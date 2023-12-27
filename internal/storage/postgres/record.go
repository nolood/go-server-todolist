package postgres

type Record struct {
	ID          int `json:"id"`
	ArticleID   int
	Article     Article `pg:"rel:has-one" json:"article"`
	Description string  `json:"description"`
}
