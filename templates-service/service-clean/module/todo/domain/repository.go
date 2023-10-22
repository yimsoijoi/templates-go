package domain

import "context"

type Repository interface {
	CreateTodo(context.Context, Todo) error
}
