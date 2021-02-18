package db

import (
	"context"
	"database/sql"
	"github.com/bramalho/go-cqrs/model"
	_ "github.com/lib/pq"
)

type PostgresRepository struct {
	db *sql.DB
}

func NewPostgres(url string) (*PostgresRepository, error) {
	db, err := sql.Open("postgres", url)
	if err != nil {
		return nil, err
	}

	return &PostgresRepository{
		db,
	}, nil
}

func (r *PostgresRepository) Close() {
	r.db.Close()
}

func (r *PostgresRepository) InsertTodo(ctx context.Context, todo model.Todo) error {
	_, err := r.db.Exec(
		"INSERT INTO todos(id, body, created_at) VALUES($1, $2, $3)",
		todo.ID, todo.Body, todo.CreatedAt)

	return err
}

func (r *PostgresRepository) ListTodos(ctx context.Context, skip uint64, take uint64) ([]model.Todo, error) {
	rows, err := r.db.Query(
		"SELECT * FROM todos ORDER BY id DESC OFFSET $1 LIMIT $2",
		skip, take)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	todos := []model.Todo{}
	for rows.Next() {
		todo := model.Todo{}
		if err = rows.Scan(&todo.ID, &todo.Body, &todo.CreatedAt); err == nil {
			todos = append(todos, todo)
		}
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return todos, nil
}
