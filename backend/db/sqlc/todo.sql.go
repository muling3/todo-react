// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0
// source: todo.sql

package db

import (
	"context"
	"database/sql"
	"log"
)

const createTodo = `-- name: CreateTodo :exec
INSERT INTO todos (
  user_id, title, body, priority, due_date
) VALUES (
  ?, ?, ?, ?, ?
)
`

type CreateTodoParams struct {
	UserID   sql.NullInt32 `json:"user_id"`
	Title    string        `json:"title"`
	Body     string        `json:"body"`
	Priority string        `json:"priority"`
	DueDate  sql.NullTime  `json:"due_date"`
}

func (q *Queries) CreateTodo(ctx context.Context, arg CreateTodoParams) error {
	_, err := q.db.ExecContext(ctx, createTodo,
		arg.UserID,
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

const deleteUserTodo = `-- name: DeleteUserTodo :exec
DELETE FROM todos
WHERE user_id = ? AND id = ?
`

type DeleteUserTodoParams struct {
	UserID sql.NullInt32 `json:"user_id"`
	ID     int32         `json:"id"`
}

func (q *Queries) DeleteUserTodo(ctx context.Context, arg DeleteUserTodoParams) error {
	_, err := q.db.ExecContext(ctx, deleteUserTodo, arg.UserID, arg.ID)
	return err
}

const getTodo = `-- name: GetTodo :one
SELECT id, user_id, title, body, priority, created_at, due_date FROM todos
WHERE id = ? LIMIT 1
`

func (q *Queries) GetTodo(ctx context.Context, id int32) (Todo, error) {
	row := q.db.QueryRowContext(ctx, getTodo, id)
	var i Todo
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.Title,
		&i.Body,
		&i.Priority,
		&i.CreatedAt,
		&i.DueDate,
	)
	return i, err
}

const getUuserTodo = `-- name: GetUuserTodo :one
SELECT id, user_id, title, body, priority, created_at, due_date FROM todos
WHERE user_id = ? AND id = ? LIMIT 1
`

type GetUuserTodoParams struct {
	UserID sql.NullInt32 `json:"user_id"`
	ID     int32         `json:"id"`
}

func (q *Queries) GetUuserTodo(ctx context.Context, arg GetUuserTodoParams) (Todo, error) {
	row := q.db.QueryRowContext(ctx, getUuserTodo, arg.UserID, arg.ID)
	var i Todo
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.Title,
		&i.Body,
		&i.Priority,
		&i.CreatedAt,
		&i.DueDate,
	)
	return i, err
}

const listTodos = `-- name: ListTodos :many
SELECT id, user_id, title, body, priority, created_at, due_date FROM todos
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
			&i.UserID,
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

const listUserTodos = `-- name: ListUserTodos :many
SELECT id, user_id, title, body, priority, created_at, due_date FROM todos
WHERE user_id = ?
ORDER BY id
`

func (q *Queries) ListUserTodos(ctx context.Context, userID sql.NullInt32) ([]Todo, error) {
	log.Printf(" CONTEXT %v AND USER_ID %v", listUserTodos, userID)
	rows, err := q.db.QueryContext(ctx, listUserTodos, userID)
	// rows, err := q.db.QueryContext(ctx, "select id, user_id, title, body, priority, created_at, due_date from todos where user_id = ?", 1)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Todo
	for rows.Next() {
		var i Todo
		if err := rows.Scan(
			&i.ID,
			&i.UserID,
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

const updateUserTodo = `-- name: UpdateUserTodo :exec
UPDATE todos SET body = ?, priority = ? WHERE user_id = ? AND id = ?
`

type UpdateUserTodoParams struct {
	Body     string        `json:"body"`
	Priority string        `json:"priority"`
	UserID   sql.NullInt32 `json:"user_id"`
	ID       int32         `json:"id"`
}

func (q *Queries) UpdateUserTodo(ctx context.Context, arg UpdateUserTodoParams) error {
	_, err := q.db.ExecContext(ctx, updateUserTodo,
		arg.Body,
		arg.Priority,
		arg.UserID,
		arg.ID,
	)
	return err
}
