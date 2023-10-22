package database

import (
	"context"
	"example.com/service-clean/module/todo/domain"
	"gorm.io/gorm"
)

type Repository interface {
	CreateTodo(context.Context, domain.Todo) error
}

type Wrapper interface {
	Create(interface{}) *gorm.DB
	WithContext(context.Context) *gorm.DB
}

type repositoryImpl struct {
	db Wrapper
}

func New(db Wrapper) Repository {
	return &repositoryImpl{db: db}
}
