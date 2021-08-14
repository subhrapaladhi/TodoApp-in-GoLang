package model

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

var con *sql.DB

const (
	host     = "localhost"
	port     = 5432
	user     = "subpsql"
	password = "pass12sub"
	dbname   = "subpsql"
)

func Connect() *sql.DB {
	psqlUri := fmt.Sprintf("host=%s port=%d user=%s dbname=%s sslmode=disable password=%s", host, port, user, dbname, password)
	db, err := sql.Open("postgres", psqlUri)
	if err != nil {
		log.Fatal("connect error =", err)
	}
	fmt.Println("Connected to the database")
	createTableQ, errCreate := db.Query("create table TODO(name varchar(50), todo varchar(1000))")
	if errCreate != nil {
		log.Fatal("table creation error = ", errCreate)
	}
	defer createTableQ.Close()
	con = db
	return db
}
