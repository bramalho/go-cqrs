package main

import (
	"time"
)

const (
	KindTodoCreated = iota + 1
)

type TodoCreatedMessage struct {
	Kind      uint32    `json:"kind"`
	ID        string    `json:"id"`
	Body      string    `json:"body"`
	CreatedAt time.Time `json:"created_at"`
}

func newTodoCreatedMessage(id string, body string, createdAt time.Time) *TodoCreatedMessage {
	return &TodoCreatedMessage{
		Kind:      KindTodoCreated,
		ID:        id,
		Body:      body,
		CreatedAt: createdAt,
	}
}
