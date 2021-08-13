package main

import (
	"TodoApp/controller"
	"TodoApp/model"
	"fmt"
	"log"
	"net/http"
)

func main() {
	mux := controller.Register()
	db := model.Connect()
	defer db.Close()
	fmt.Println("Server started")
	log.Fatal(http.ListenAndServe("localhost:3000", mux))
}
