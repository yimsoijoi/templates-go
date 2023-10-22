package redis

import (
	"context"
	"example.com/service-clean/module/todo/domain"
	"github.com/redis/go-redis/v9"
	"time"
)

type Repository interface {
	CreateTodo(context.Context, domain.Todo) error
}

type repositoryImpl struct {
	rd Wrapper
}

func New(rd Wrapper) Repository {
	return &repositoryImpl{rd: rd}
}

type Wrapper interface {
	Set(context.Context, string, interface{}, time.Duration) *redis.StatusCmd
}
