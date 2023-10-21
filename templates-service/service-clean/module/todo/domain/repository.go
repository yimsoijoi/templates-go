package domain

type Repository interface {
	CreateTodo(todo Todo) error
}
