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

func createDefaultRecordTypes() {
	recordTypes := []RecordType{
		{Value: "income"},
		{Value: "expense"},
	}

	var count int64

	query := Db.Table("record_types")
	query.Count(&count)

	if count == 0 {
		Db.Create(recordTypes)
	}

}
