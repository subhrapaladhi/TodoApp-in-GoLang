package main

import (
	"TodoApp/structs"
	"encoding/json"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/ping", func(rw http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			data := structs.Response{
				Code: http.StatusOK,
				Body: "pong",
			}
			json.NewEncoder(rw).Encode(data)
		}
	})

	http.ListenAndServe("localhost:3000", mux)
}
