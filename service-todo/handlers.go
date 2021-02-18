package main

import (
	"github.com/bramalho/go-cqrs/db"
	"github.com/bramalho/go-cqrs/event"
	"github.com/bramalho/go-cqrs/model"
	"github.com/bramalho/go-cqrs/util"
	"github.com/segmentio/ksuid"
	"log"
	"net/http"
	"time"
	"html/template"
)

func createTodoHandler(w http.ResponseWriter, r *http.Request) {
	type response struct {
		ID string `json:"id"`
	}

	ctx := r.Context()

	// Read parameters
	body := template.HTMLEscapeString(r.FormValue("body"))
	if len(body) < 1 || len(body) > 140 {
		util.ResponseError(w, http.StatusBadRequest, "Invalid body")
		return
	}

	// Create todo
	createdAt := time.Now().UTC()
	id, err := ksuid.NewRandomWithTime(createdAt)
	if err != nil {
		util.ResponseError(w, http.StatusInternalServerError, "Failed to create todo")
		return
	}
	todo := model.Todo{
		ID:        id.String(),
		Body:      body,
		CreatedAt: createdAt,
	}
	if err := db.InsertTodo(ctx, todo); err != nil {
		log.Println(err)
		util.ResponseError(w, http.StatusInternalServerError, "Failed to create todo")
		return
	}

	// Publish event
	if err := event.PublishTodoCreated(todo); err != nil {
		log.Println(err)
	}

	// Return new todo
	util.ResponseOk(w, response{ID: todo.ID})
}
