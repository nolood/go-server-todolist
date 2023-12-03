package handlers

import (
	"net/http"
)

func GetAllTasks(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("You get all tasks !!!"))
}
