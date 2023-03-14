// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.15.0

package db

import (
	"database/sql"
)

type Todo struct {
	ID        int32        `json:"id"`
	Title     string       `json:"title"`
	Body      string       `json:"body"`
	Priority  string       `json:"priority"`
	CreatedAt sql.NullTime `json:"created_at"`
	DueDate   sql.NullTime `json:"due_date"`
}