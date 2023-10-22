package create_todo

import (
	"errors"
	"example.com/service-clean/module/todo/common"
)

var (
	errCreateTodoUseCase = errors.New("failed to create todo  usecase")
)

func (self createTodoImpl) wrapError(err error) error {
	return common.WrapError(errCreateTodoUseCase.Error(), err)
}
