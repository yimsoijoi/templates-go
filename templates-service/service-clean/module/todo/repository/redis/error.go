package redis

import (
	"errors"
	"example.com/service-clean/module/todo/common"
)

var (
	errCreateTodoRedisRepository = errors.New("failed to create todo redis repository")
)

func wrapError(errContext string, err error) error {
	return common.WrapError(errContext, err)
}
