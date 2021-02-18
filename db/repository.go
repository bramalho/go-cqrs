package db

import (
	"context"
	"github.com/bramalho/go-cqrs/model"
)

type Repository interface {
	Close()
	InsertTodo(ctx context.Context, todo model.Todo) error
	ListTodos(ctx context.Context, skip uint64, take uint64) ([]model.Todo, error)
}

var impl Repository

func SetRepository(repository Repository) {
	impl = repository
}

func Close() {
	impl.Close()
}

func InsertTodo(ctx context.Context, todo model.Todo) error {
	return impl.InsertTodo(ctx, todo)
}

func ListTodos(ctx context.Context, skip uint64, take uint64) ([]model.Todo, error) {
	return impl.ListTodos(ctx, skip, take)
}
