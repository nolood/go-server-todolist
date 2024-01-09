package handlers

import (
	"go-server/internal/config"
	"go-server/internal/storage/postgres"
	"log"
	"net/http"
)

func CreateRecord(w http.ResponseWriter, r *http.Request) {
	userId, err := getUserId(r)
	if err != nil {
		config.Logger.Error(err.Error())
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	var record postgres.Record

	log.Println(userId)

	err = fromBody(r.Body, &record)
	if err != nil {
		config.Logger.Error(err.Error())
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}
}
