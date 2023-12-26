package handlers

import (
	"encoding/json"
	"io"
	"log"

	"github.com/go-chi/render"
)

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
