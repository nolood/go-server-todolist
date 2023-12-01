package tasks

import "net/http"

func SendMessage(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World!"))
}

func GetAllTasks(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("You get all tasks !!!"))
}
