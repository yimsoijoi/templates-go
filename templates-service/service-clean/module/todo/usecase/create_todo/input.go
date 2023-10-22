package create_todo

import (
	"example.com/service-clean/module/todo/domain"
	"github.com/google/uuid"
	"time"
)

func (self createTodoImpl) prepare(input domain.TodoCreateInput) domain.Todo {
	return domain.Todo{
		ID:          uuid.NewString(),
		Title:       input.Title,
		Description: input.Description,
		Status:      input.Status,
		TodoDate:    input.TodoDate,
		CreatedAt:   time.Now(),
	}
}
