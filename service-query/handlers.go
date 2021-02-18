package main

import (
	"context"
	"github.com/bramalho/go-cqrs/db"
	"github.com/bramalho/go-cqrs/event"
	"github.com/bramalho/go-cqrs/model"
	"github.com/bramalho/go-cqrs/search"
	"github.com/bramalho/go-cqrs/util"
	"log"
	"net/http"
	"strconv"
)

func onTodoCreated(m event.TodoCreatedMessage) {
	todo := model.Todo{
		ID:        m.ID,
		Body:      m.Body,
		CreatedAt: m.CreatedAt,
	}
	if err := search.InsertTodo(context.Background(), todo); err != nil {
		log.Println(err)
	}
}

func listTodosHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	var err error

	// Read parameters
	skip := uint64(0)
	skipStr := r.FormValue("skip")
	take := uint64(100)
	takeStr := r.FormValue("take")
	if len(skipStr) != 0 {
		skip, err = strconv.ParseUint(skipStr, 10, 64)
		if err != nil {
			util.ResponseError(w, http.StatusBadRequest, "Invalid skip parameter")
			return
		}
	}
	if len(takeStr) != 0 {
		take, err = strconv.ParseUint(takeStr, 10, 64)
		if err != nil {
			util.ResponseError(w, http.StatusBadRequest, "Invalid take parameter")
			return
		}
	}

	// Fetch todos
	todos, err := db.ListTodos(ctx, skip, take)
	if err != nil {
		log.Println(err)
		util.ResponseError(w, http.StatusInternalServerError, "Could not fetch todos")
		return
	}

	util.ResponseOk(w, todos)
}

func searchTodosHandler(w http.ResponseWriter, r *http.Request) {
	var err error
	ctx := r.Context()

	// Read parameters
	query := r.FormValue("query")
	if len(query) == 0 {
		util.ResponseError(w, http.StatusBadRequest, "Missing query parameter")
		return
	}
	skip := uint64(0)
	skipStr := r.FormValue("skip")
	take := uint64(100)
	takeStr := r.FormValue("take")
	if len(skipStr) != 0 {
		skip, err = strconv.ParseUint(skipStr, 10, 64)
		if err != nil {
			util.ResponseError(w, http.StatusBadRequest, "Invalid skip parameter")
			return
		}
	}
	if len(takeStr) != 0 {
		take, err = strconv.ParseUint(takeStr, 10, 64)
		if err != nil {
			util.ResponseError(w, http.StatusBadRequest, "Invalid take parameter")
			return
		}
	}

	// Search todos
	todos, err := search.SearchTodos(ctx, query, skip, take)
	if err != nil {
		log.Println(err)
		util.ResponseOk(w, []model.Todo{})
		return
	}

	util.ResponseOk(w, todos)
}
