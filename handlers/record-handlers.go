package handlers

import (
	"go-server/internal/config"
	"go-server/internal/storage/postgres"
	"net/http"
)

type CreateRecordParam struct {
	ID           uint64              `json:"id" gorm:"autoIncrement"`
	Article      postgres.Article    `json:"article"`
	ArticleID    uint64              `json:"article_id"`
	Description  string              `json:"description"`
	BillID       uint64              `json:"bill_id"`
	Amount       int                 `json:"amount"`
	RecordType   postgres.RecordType `json:"type"`
	RecordTypeID uint64              `json:"type_id"`
	Date         string              `json:"date"`
}

func CreateRecord(w http.ResponseWriter, r *http.Request) {
	userId, err := getUserId(r)
	if err != nil {
		config.Logger.Error(err.Error())
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	var record postgres.Record

	err = fromBody(r.Body, &record)
	if err != nil {
		config.Logger.Error(err.Error())
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	var recordType postgres.RecordType

	query := postgres.Db.Table("record_types")
	result := query.Where("id = ?", record.RecordTypeID).Find(&recordType)
	if result.Error != nil {
		config.Logger.Error(result.Error.Error())
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	var bill postgres.Bill

	query = postgres.Db.Table("bills")
	result = query.Where("user_id = ?", userId).Find(&bill)
	if result.Error != nil {
		config.Logger.Error(result.Error.Error())
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	query = postgres.Db.Table("records")
	query.Create(&record)

	if record.RecordTypeID == 1 {
		bill.Balance += float32(record.Amount)
	} else if record.RecordTypeID == 2 {
		bill.Balance -= float32(record.Amount)
	}

	postgres.Db.Save(&bill)

	var recordResponse CreateRecordParam

	query = postgres.Db.Table("records")
	query.Where("id = ?", record.ID).Preload("RecordType").Preload("Article").Find(&recordResponse)
	if query.Error != nil {
		config.Logger.Error(query.Error.Error())
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	w.Write(toJson(recordResponse))
}
