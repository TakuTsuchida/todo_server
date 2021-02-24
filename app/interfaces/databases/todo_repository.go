package databases

import "app/domains"

type TodoRepository struct {
	SqlHandler
}

func (db *TodoRepository) SelectAll() (todos domains.Todos) {
	db.FindAll(&todos)
	return
}
