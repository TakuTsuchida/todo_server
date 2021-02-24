package interactors

import (
	"app/domains"
	"app/usecases/repositories"
)

type TodoInteractor struct {
	TodoRepository repositories.TodoRepository
}

func (interactor *TodoInteractor) Todos() (todos domains.Todos) {
	interactor.TodoRepository.SelectAll()
	return
}
