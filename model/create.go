package model

import "log"

func CreateTodo(name, todo string) error {
	insertQ, err := con.Query("INSERT INTO TODO VALUES($1, $2)", name, todo)
	if err != nil {
		log.Fatal("create todo = ", err)
		return err
	}
	defer insertQ.Close()
	return nil
}
