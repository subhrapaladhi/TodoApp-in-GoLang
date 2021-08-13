package model

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

var con *sql.DB

func Connect() *sql.DB {
	db, err := sql.Open("postgres", "subpsql:pass12sub@(tcp:localhost:5432)")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected to the database")
	con = db
	return db
}
