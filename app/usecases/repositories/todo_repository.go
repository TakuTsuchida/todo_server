package repositories

import "app/domains"

type TodoRepository interface {
	SelectAll() domains.Todos
}
