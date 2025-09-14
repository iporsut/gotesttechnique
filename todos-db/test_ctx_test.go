package main

import (
	"database/sql"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/testcontainers/testcontainers-go/modules/postgres"
)

type testContext struct {
	Repo *TodoRepo
	t    *testing.T
}

func NewTestContext(t *testing.T) *testContext {
	t.Helper()

	ctx := t.Context()
	postgresContainer, err := postgres.Run(ctx,
		"postgres:16-alpine",
		postgres.WithDatabase("test"),
		postgres.WithUsername("user"),
		postgres.WithPassword("password"),
		postgres.BasicWaitStrategies(),
	)
	require.NoError(t, err)

	connStr, err := postgresContainer.ConnectionString(ctx, "sslmode=disable")
	require.NoError(t, err)

	db, err := sql.Open("postgres", connStr)
	require.NoError(t, err)
	t.Cleanup(func() {
		db.Close()
	})

	err = migrateUp(db)
	require.NoError(t, err)
	repo := NewTodoRepo(db)

	return &testContext{Repo: repo, t: t}
}

func (tc *testContext) CreateTodo(title string, completed bool) *Todo {
	tc.t.Helper()
	todo := &Todo{Title: title, Completed: completed}
	savedTodo, err := tc.Repo.Create(todo)
	require.NoError(tc.t, err)
	return savedTodo
}
