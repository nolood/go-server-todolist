package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/go-chi/render"
	"io"
	"log"
	"net/http"
	"reflect"
	"strconv"
	"strings"
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

func validateQueryParams(r *http.Request, target interface{}) error {
	targetType := reflect.TypeOf(target).Elem()
	targetValue := reflect.ValueOf(target).Elem()

	for i := 0; i < targetType.NumField(); i++ {
		field := targetType.Field(i)
		tag := field.Tag.Get("json")
		paramValue := r.URL.Query().Get(tag)

		err := setField(paramValue, targetValue.Field(i))
		if err != nil {
			return err
		}
	}

	return nil
}

func setField(value string, field reflect.Value) error {
	switch field.Kind() {
	case reflect.Int:
		if value == "" {
			field.SetInt(0)
			return nil
		}
		intValue, err := strconv.Atoi(value)
		if err != nil {
			return err
		}
		field.SetInt(int64(intValue))
	case reflect.String:
		field.SetString(value)
	case reflect.Array, reflect.Slice:
		if value == "" {
			field.Set(reflect.MakeSlice(field.Type(), 0, 0))
			return nil
		}
		numStrs := strings.Split(value, ",")
		slice := reflect.MakeSlice(field.Type(), len(numStrs), len(numStrs))

		for i, numStr := range numStrs {
			num, err := strconv.Atoi(strings.TrimSpace(numStr))
			if err != nil {
				return err
			}
			slice.Index(i).SetInt(int64(num))
		}

		field.Set(slice)
	default:
		return fmt.Errorf("Unsupported type: %v", field.Kind())
	}
	return nil
}
