package event

import (
	"github.com/bramalho/go-cqrs/model"
)

type EventStore interface {
	Close()
	PublishTodoCreated(todo model.Todo) error
	SubscribeTodoCreated() (<-chan TodoCreatedMessage, error)
	OnTodoCreated(f func(TodoCreatedMessage)) error
}

var impl EventStore

func SetEventStore(es EventStore) {
	impl = es
}

func Close() {
	impl.Close()
}

func PublishTodoCreated(todo model.Todo) error {
	return impl.PublishTodoCreated(todo)
}

func SubscribeTodoCreated() (<-chan TodoCreatedMessage, error) {
	return impl.SubscribeTodoCreated()
}

func OnTodoCreated(f func(TodoCreatedMessage)) error {
	return impl.OnTodoCreated(f)
}
