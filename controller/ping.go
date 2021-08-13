package controller

import (
	"TodoApp/views"
	"encoding/json"
	"net/http"
)

func ping() http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			data := views.Response{
				Code: http.StatusOK,
				Body: "pong",
			}
			rw.WriteHeader(http.StatusOK)
			json.NewEncoder(rw).Encode(data)
		}
	}
}
