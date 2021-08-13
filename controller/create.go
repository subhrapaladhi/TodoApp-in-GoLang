package controller

import (
	"TodoApp/model"
	"TodoApp/views"
	"encoding/json"
	"net/http"
)

func create() http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			data := views.PostRequest{}
			json.NewDecoder(r.Body).Decode(&data)
			if err := model.CreateTodo(data.Name, data.Todo); err != nil {
				rw.Write([]byte("Error creating todo"))
				return
			}
			rw.WriteHeader(http.StatusCreated)
			json.NewEncoder(rw).Encode(data)
		}
	}
}
