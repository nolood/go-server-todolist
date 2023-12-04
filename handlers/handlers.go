package handlers

import (
	"encoding/json"
	"github.com/go-chi/render"
	"github.com/go-pg/pg/v10"
	"go-server/internal/storage/postgres"
	"io"
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

func fromBody(body io.ReadCloser, model interface{}) error {
	err := render.DecodeJSON(body, &model)

	return err
}
