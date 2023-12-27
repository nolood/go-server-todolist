package postgres

type ArticleType struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type Article struct {
	ID            int         `json:"id"`
	Icon          string      `json:"icon"`
	ArticleTypeID int         `pg:"alias:type_id"`
	ArticleType   ArticleType `pg:"rel:has-one,alias:type" json:"type"`
}
