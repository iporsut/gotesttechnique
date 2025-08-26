package main

import (
	"context"
	"database/sql"
	"testing"

	_ "github.com/lib/pq"
	"github.com/pressly/goose/v3"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/testcontainers/testcontainers-go/modules/postgres"
)

func migrateUp(db *sql.DB) error {
	err := goose.SetDialect("postgres")
	if err != nil {
		return err
	}

	return goose.Up(db, "./migrations")
}

func TestRepoSave(t *testing.T) {
	ctx := context.Background()
	postgresContainer, err := postgres.Run(ctx,
		"postgres:16-alpine",
		postgres.WithDatabase("test"),
		postgres.WithUsername("user"),
		postgres.WithPassword("password"),
		postgres.BasicWaitStrategies(),
	)

	connStr, err := postgresContainer.ConnectionString(ctx, "sslmode=disable")
	require.NoError(t, err)

	db, err := sql.Open("postgres", connStr)
	require.NoError(t, err)
	defer db.Close()

	err = migrateUp(db)
	require.NoError(t, err)

	repo := NewTodoRepo(db)
	todo := &Todo{Title: "Test Todo", Completed: false}
	savedTodo, err := repo.Create(todo)
	require.NoError(t, err)

	assert.NotZero(t, savedTodo.ID)
	assert.Equal(t, "Test Todo", savedTodo.Title)
	assert.False(t, savedTodo.Completed)
}
