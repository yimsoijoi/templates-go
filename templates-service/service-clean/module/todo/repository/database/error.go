package database

import (
	"errors"
	"example.com/service-clean/module/todo/common"
)

var (
	errCreateTodoDBRepository = errors.New("failed to create todo db repository")
)

func wrapError(errContext string, err error) error {
	return common.WrapError(errContext, err)
}
