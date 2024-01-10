package postgres

type RecordType struct {
	Model
	Value string `json:"value"`
}

type Record struct {
	Model
	Article      Article    `json:"article"`
	ArticleID    uint64     `json:"article_id"`
	Description  string     `json:"description"`
	Bill         Bill       `json:"bill"`
	BillID       uint64     `json:"bill_id"`
	Amount       int        `json:"amount"`
	RecordType   RecordType `json:"type"`
	RecordTypeID uint64     `json:"type_id"`
	Date         string     `json:"date"`
}
