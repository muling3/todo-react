package db

import (
	"context"
	"database/sql"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestCreateTodo(t *testing.T){
	createTodoParams := CreateTodoParams{
		Title: "This is dummy title",
		Body: "This is dummy title body which ends here",
		DueDate: sql.NullTime{
			Time: time.Now().AddDate(0, 0, 2),
			Valid: true,
		},
		Priority: "LOW",
	}

	err := testQueries.CreateTodo(context.Background(), createTodoParams)

	require.Empty(t, err)
}

func TestGetTodos(t *testing.T){
	todos, err := testQueries.ListTodos(context.Background())

	require.Empty(t, err)
	require.Len(t, todos, 1)
}

func TestGetTodo(t *testing.T){ 
	todo, err := testQueries.GetTodo(context.Background(), 1)

	require.Empty(t, err)
	require.Equal(t, int32(1), todo.ID)
}

func TestGetTodoToThrowError(t *testing.T){
	todo, err := testQueries.GetTodo(context.Background(), 10)

	require.NotEmpty(t, err)
	require.Empty(t, todo)
	require.EqualError(t, sql.ErrNoRows, err.Error())
}


func TestUpdateTodo(t *testing.T){
	updateTodoParams := UpdateTodoParams{
		Body: "This is updated title body which ends here",
		Priority: "LOW",
		ID: 3,
	}

	err := testQueries.UpdateTodo(context.Background(), updateTodoParams)
	require.Empty(t, err)
}

func TestDeleteTodo(t *testing.T){
	err := testQueries.DeleteTodo(context.Background(), 1)

	require.Empty(t, err)
}