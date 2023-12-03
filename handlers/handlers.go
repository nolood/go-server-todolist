package handlers

import (
	"encoding/json"
	"github.com/go-pg/pg/v10"
	"go-server/internal/storage/postgres"
	"log"
)

var (
	db *pg.DB
)

func init() {
	db = postgres.ConnectDb()
}

func toJson(model any) []byte {
	result, err := json.Marshal(model)
	if err != nil {
		log.Fatal(err)
	}

	return result
}

func fromJson(model any) {
}
