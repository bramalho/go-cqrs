package event

import (
	"time"
)

type Message interface {
	Key() string
}

type TodoCreatedMessage struct {
	ID        string
	Body      string
	CreatedAt time.Time
}

func (m *TodoCreatedMessage) Key() string {
	return "todo.created"
}
