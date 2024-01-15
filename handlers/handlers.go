package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/go-chi/render"
	"io"
	"log"
	"net/http"
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

func getUserId(r *http.Request) (uint64, error) {
	userId, ok := r.Context().Value("user_id").(uint64)
	if !ok {
		return userId, fmt.Errorf("user ID not found")
	}
	return userId, nil
}
