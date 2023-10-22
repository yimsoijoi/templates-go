package create_todo

import (
	"example.com/service-clean/module/todo/common"
	"example.com/service-clean/module/todo/domain"
)

func (self createTodoImpl) validateInput(input domain.TodoCreateInput) error {
	return self.wrapError(common.ValidateStruct(input))
}
