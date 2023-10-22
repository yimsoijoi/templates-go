package create_todo

import (
	"context"
	"example.com/service-clean/module/todo/domain"
)

type CreateTodoUseCase interface {
	CreateTodo(ctx context.Context, entity domain.TodoCreateInput) error
}

type createTodoImpl struct {
	repo domain.Repository
}

func New(repo domain.Repository) CreateTodoUseCase {
	return &createTodoImpl{repo: repo}
}

func (self createTodoImpl) CreateTodo(ctx context.Context, entity domain.TodoCreateInput) error {
	if err := self.validateInput(entity); err != nil {
		return err
	}

	return self.wrapError(self.repo.CreateTodo(ctx, self.prepare(entity)))
}
