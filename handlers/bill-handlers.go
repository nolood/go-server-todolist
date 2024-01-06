package handlers

import (
	"go-server/internal/config"
	"go-server/internal/storage/postgres"
	"net/http"
	"time"
)

type CreateBillParam struct {
	ID      uint64  `json:"id"`
	Title   string  `json:"title"`
	Balance float64 `json:"balance"`
	UserID  uint64  `json:"user_id"`
}

func CreateBill(w http.ResponseWriter, r *http.Request) {
	userId, err := getUserId(r)
	if err != nil {
		config.Logger.Error(err.Error())
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	var bill CreateBillParam

	bill.UserID = userId

	err = fromBody(r.Body, &bill)
	if err != nil {
		config.Logger.Error(err.Error())
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	query := postgres.Db.Table("bills")

	query = query.Create(&bill)

	w.Write(toJson(bill))
}

type GetBillParam struct {
	Id        uint64    `json:"id"`
	Title     string    `json:"title"`
	Balance   float64   `json:"balance"`
	UserId    uint64    `json:"user_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func GetAllBills(w http.ResponseWriter, r *http.Request) {
	userId, err := getUserId(r)
	if err != nil {
		config.Logger.Error(err.Error())
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	var bills []GetBillParam

	query := postgres.Db.Table("bills")

	query = query.Where("user_id = ?", userId)

	query.Find(&bills)

	w.Write(toJson(bills))
}
