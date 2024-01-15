package handlers

import (
	"fmt"
	"github.com/go-chi/chi/v5"
	"go-server/internal/config"
	"go-server/internal/storage/postgres"
	"gorm.io/gorm"
	"net/http"
	"strconv"
	"time"
)

type CreateBillParam struct {
	gorm.Model
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

	if len(bill.Title) > 30 {
		http.Error(w, "Title should be less than 30", http.StatusBadRequest)
		return
	}

	query := postgres.Db.Table("bills")

	var count int64

	query.Where("user_id = ?", userId).Count(&count)

	if count >= 6 {
		http.Error(w, "You can't create more than 6 bills", http.StatusBadRequest)
		return
	}

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

	query = query.Where("user_id = ?", userId).Order("created_at")

	query.Find(&bills)

	w.Write(toJson(bills))
}

func GetBill(w http.ResponseWriter, r *http.Request) {
	userId, err := getUserId(r)
	if err != nil {
		config.Logger.Error(err.Error())
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	billIdParam := chi.URLParam(r, "id")

	billID, err := strconv.Atoi(billIdParam)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	var bill GetBillParam

	query := postgres.Db.Table("bills")

	query = query.Where("user_id = ?", userId)

	query = query.Where("id = ?", billID)

	query.Find(&bill)
	if query.Error != nil {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	var records []postgres.Record

	query = postgres.Db.Table("records")

	query = query.Where("bill_id = ?", billID).Preload("RecordType")

	query.Find(&records)
	if query.Error != nil {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	type Response struct {
		GetBillParam
		TotalIncome  int `json:"total_income"`
		TotalExpense int `json:"total_expense"`
	}

	response := Response{
		GetBillParam: bill,
		TotalIncome:  0.00,
		TotalExpense: 0.00,
	}

	for _, record := range records {
		switch record.RecordType.Value {
		case "income":
			response.TotalIncome += record.Amount
		case "expense":
			response.TotalExpense += record.Amount
		}
	}

	w.Write(toJson(response))
}

func changeBillBalance(billID uint64, userID uint64, amount float32, Type uint64) error {
	var bill postgres.Bill

	query := postgres.Db.Table("bills")
	query.Where("id = ?", billID).Where("user_id = ?", userID).Find(&bill)
	if query.Error != nil {
		return fmt.Errorf("can't find bill")
	}

	if Type == 1 {
		bill.Balance += amount
	} else {
		bill.Balance -= amount
	}

	postgres.Db.Save(&bill)

	return nil
}
