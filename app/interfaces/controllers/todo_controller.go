package controllers

import (
	"app/domains"
	"app/interfaces/databases"
	"app/usecases/interactors"
)

type TodoController struct {
	Interactor interactors.TodoInteractor
}

func NewTodoController(sqlHandler databases.SqlHandler) *TodoController {
	return &TodoController{
		Interactor: interactors.TodoInteractor{
			TodoRepository: &databases.TodoRepository{
				SqlHandler: sqlHandler,
			},
		},
	}
}

func (controller *TodoController) GetAll() (todos domains.Todos) {
	todos = controller.Interactor.Todos()
	return
}
