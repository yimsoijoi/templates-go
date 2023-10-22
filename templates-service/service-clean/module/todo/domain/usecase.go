package domain

import "context"

type UseCase interface {
	CreateTodo(ctx context.Context, todo TodoCreateInput) error
}
