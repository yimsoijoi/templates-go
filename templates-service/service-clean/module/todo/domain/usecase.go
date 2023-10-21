package domain

type UseCase interface {
	CreateTodo(todo Todo) error
}
