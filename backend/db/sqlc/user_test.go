package db

import (
	"context"
	"database/sql"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCreateUser(t *testing.T) {
	createUserParams := CreateUserParams{
		Username: "dummyuser",
		Password: "userpass123",
	}

	err := testQueries.CreateUser(context.Background(), createUserParams)

	require.Empty(t, err)
}

func TestGetUsers(t *testing.T) {
	users, err := testQueries.ListUsers(context.Background())

	require.Empty(t, err)
	require.GreaterOrEqual(t, len(users), 1)
}

func TestGetUser(t *testing.T) {
	user, err := testQueries.GetUser(context.Background(), 3)

	require.Empty(t, err)
	require.Equal(t, int32(3), user.ID)
}

func TestGetUserToThrowError(t *testing.T) {
	user, err := testQueries.GetUser(context.Background(), 10)

	require.NotEmpty(t, err)
	require.Empty(t, user)
	require.EqualError(t, sql.ErrNoRows, err.Error())
}

func TestUpdateUser(t *testing.T) {
	updateUserParams := UpdateUserParams{
		Username: "dummyuser",
		Password: "userpass123",
		ID:       3,
	}

	err := testQueries.UpdateUser(context.Background(), updateUserParams)
	require.Empty(t, err)
}

func TestDeleteUser(t *testing.T) {
	err := testQueries.DeleteUser(context.Background(), 1)

	require.Empty(t, err)
}
