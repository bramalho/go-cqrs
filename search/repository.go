package search

import (
	"context"
	"github.com/bramalho/go-cqrs/model"
)

type Repository interface {
	Close()
	InsertTodo(ctx context.Context, todo model.Todo) error
	SearchTodos(ctx context.Context, query string, skip uint64, take uint64) ([]model.Todo, error)
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

func SearchTodos(ctx context.Context, query string, skip uint64, take uint64) ([]model.Todo, error) {
	return impl.SearchTodos(ctx, query, skip, take)
}
