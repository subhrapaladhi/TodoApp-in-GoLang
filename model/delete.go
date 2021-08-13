package model

import "log"

func DeleteTodo(name string) error {
	deleteQ, err := con.Query("delete from TODO where name=$1", name)
	if err != nil {
		log.Fatal("delete todo = ", err)
		return err
	}
	defer deleteQ.Close()
	return nil
}
