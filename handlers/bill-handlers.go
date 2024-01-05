package handlers

import (
	"go-server/internal/config"
	"go-server/internal/storage/postgres"
	"net/http"
)

func CreateBill(w http.ResponseWriter, r *http.Request) {
	userId, err := getUserId(r)
	if err != nil {
		config.Logger.Error(err.Error())
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	var bill postgres.Bill

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

func GetAllBills(w http.ResponseWriter, r *http.Request) {
	userId, err := getUserId(r)
	if err != nil {
		config.Logger.Error(err.Error())
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	var bills []postgres.Bill

	query := postgres.Db.Table("bills")

	query = query.Where("user_id = ?", userId)

	query.Find(&bills)

	w.Write(toJson(bills))
}
