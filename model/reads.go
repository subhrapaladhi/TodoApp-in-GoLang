package model

import "TodoApp/views"

func ReadAll() ([]views.PostRequest, error) {
	rows, err := con.Query("select * from TODO")

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	todos := []views.PostRequest{}
	for rows.Next() {
		data := views.PostRequest{}
		rows.Scan(&data.Name, &data.Todo)
		todos = append(todos, data)
	}

	return todos, nil
}

func ReadByName(name string) ([]views.PostRequest, error) {
	rows, err := con.Query("select * from TODO where name=$1", name)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	todos := []views.PostRequest{}
	for rows.Next() {
		data := views.PostRequest{}
		rows.Scan(&data.Name, &data.Todo)
		todos = append(todos, data)
	}

	return todos, nil
}
