package domain

import "time"

type Status string

const (
	StatusTodo       Status = "TODO"
	StatusInProgress Status = "IN_PROGRESS"
	StatusDone       Status = "DONE"
)

type Todo struct {
	ID          string
	Title       string
	Description string
	Status      Status
	TodoDate    time.Time
	CreatedAt   time.Time
}

func NewTodo(
	id string,
	title string,
	description string,
	status Status,
	todoDate time.Time,
	createdAt time.Time,
) *Todo {
	return &Todo{
		ID:          id,
		Title:       title,
		Description: description,
		Status:      status,
		TodoDate:    todoDate,
		CreatedAt:   createdAt,
	}
}
