package usecase

import (
	"context"
	"example.com/service-clean/module/todo/domain"
)

func (usecase useCaseImpl) CreateTodo(ctx context.Context, entity domain.TodoCreateInput) error {
	return usecase.CreateTodoUseCase.CreateTodo(ctx, entity)
}
