package controller

import (
	"TodoApp/model"
	"TodoApp/views"
	"encoding/json"
	"net/http"
)

func crud() http.HandlerFunc {
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
		} else if r.Method == http.MethodGet {
			name := r.URL.Query().Get("name")

			var data []views.PostRequest
			var err error

			if name != "" {
				data, err = model.ReadByName(name)
			} else {
				data, err = model.ReadAll()
			}

			if err != nil {
				rw.Write([]byte(err.Error()))
			}

			rw.WriteHeader(http.StatusOK)
			json.NewEncoder(rw).Encode(data)
		} else if r.Method == http.MethodDelete {
			name := r.URL.Path[1:]

			if err := model.DeleteTodo(name); err != nil {
				rw.Write([]byte("Error creating todo"))
				return
			}

			rw.WriteHeader(http.StatusOK)
			json.NewEncoder(rw).Encode(struct {
				Status string `json:"status"`
			}{"Item deleted"})
		}
	}
}
