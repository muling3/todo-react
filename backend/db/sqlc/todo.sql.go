// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.15.0
// source: todo.sql

package db

import (
	"context"
	"database/sql"
)

const createTodo = `-- name: CreateTodo :exec
INSERT INTO todos (
  title, body, priority, due_date
) VALUES (
  ?, ?, ?, ?
)
`

type CreateTodoParams struct {
	Title    string       `json:"title"`
	Body     string       `json:"body"`
	Priority string       `json:"priority"`
	DueDate  sql.NullTime `json:"due_date"`
}

func (q *Queries) CreateTodo(ctx context.Context, arg CreateTodoParams) error {
	_, err := q.db.ExecContext(ctx, createTodo,
		arg.Title,
		arg.Body,
		arg.Priority,
		arg.DueDate,
	)
	return err
}

const deleteTodo = `-- name: DeleteTodo :exec
DELETE FROM todos
WHERE id = ?
`

func (q *Queries) DeleteTodo(ctx context.Context, id int32) error {
	_, err := q.db.ExecContext(ctx, deleteTodo, id)
	return err
}

const getTodo = `-- name: GetTodo :one
SELECT id, title, body, priority, created_at, due_date FROM todos
WHERE id = ? LIMIT 1
`

func (q *Queries) GetTodo(ctx context.Context, id int32) (Todo, error) {
	row := q.db.QueryRowContext(ctx, getTodo, id)
	var i Todo
	err := row.Scan(
		&i.ID,
		&i.Title,
		&i.Body,
		&i.Priority,
		&i.CreatedAt,
		&i.DueDate,
	)
	return i, err
}

const listTodos = `-- name: ListTodos :many
SELECT id, title, body, priority, created_at, due_date FROM todos
ORDER BY id
`

func (q *Queries) ListTodos(ctx context.Context) ([]Todo, error) {
	rows, err := q.db.QueryContext(ctx, listTodos)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Todo
	for rows.Next() {
		var i Todo
		if err := rows.Scan(
			&i.ID,
			&i.Title,
			&i.Body,
			&i.Priority,
			&i.CreatedAt,
			&i.DueDate,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateTodo = `-- name: UpdateTodo :exec
UPDATE todos SET body = ?, priority = ? WHERE id = ?
`

type UpdateTodoParams struct {
	Body     string `json:"body"`
	Priority string `json:"priority"`
	ID       int32  `json:"id"`
}

func (q *Queries) UpdateTodo(ctx context.Context, arg UpdateTodoParams) error {
	_, err := q.db.ExecContext(ctx, updateTodo, arg.Body, arg.Priority, arg.ID)
	return err
}
