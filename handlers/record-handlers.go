package handlers

import (
	"github.com/go-chi/chi/v5"
	"go-server/internal/config"
	"go-server/internal/storage/postgres"
	"net/http"
	"strconv"
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

	err = changeBillBalance(record.BillID, userId, float32(record.Amount), record.RecordTypeID)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	query = postgres.Db.Table("records")
	query.Create(&record)

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

type RecordResponse struct {
	postgres.Model
	Amount       int                 `json:"amount"`
	Description  string              `json:"description"`
	Article      postgres.Article    `json:"article"`
	ArticleID    uint64              `json:"article_id"`
	BillID       uint64              `json:"bill_id"`
	RecordType   postgres.RecordType `json:"type"`
	RecordTypeID uint64              `json:"type_id"`
	Date         string              `json:"date"`
}

func GetRecordsByBillId(w http.ResponseWriter, r *http.Request) {
	billIdParam := chi.URLParam(r, "billId")

	billID, err := strconv.Atoi(billIdParam)
	if err != nil {
		config.Logger.Error(err.Error())
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	var records []RecordResponse

	query := postgres.Db.Table("records")
	query = query.Where("bill_id = ?", billID)
	query.Preload("Article").Preload("RecordType").Find(&records)
	if query.Error != nil {
		config.Logger.Error(query.Error.Error())
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	w.Write(toJson(records))
}
