package redis

import (
	"context"
	"example.com/service-clean/module/todo/domain"
)

func (self repositoryImpl) CreateTodo(ctx context.Context, entity domain.Todo) error {
	return wrapError(errCreateTodoRedisRepository.Error(), self.rd.Set(ctx, entity.ID, entity, -1).Err())
}
