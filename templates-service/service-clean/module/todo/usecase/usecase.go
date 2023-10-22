package usecase

import (
	"example.com/service-clean/module/todo/domain"
	"example.com/service-clean/module/todo/usecase/create_todo"
)

type useCaseImpl struct {
	create_todo.CreateTodoUseCase
}

func New(repo domain.Repository) domain.UseCase {
	return &useCaseImpl{
		CreateTodoUseCase: create_todo.New(repo),
	}
}
