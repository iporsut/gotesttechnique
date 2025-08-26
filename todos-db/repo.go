package main

import "database/sql"

type TodoRepo struct {
	db *sql.DB
}

func NewTodoRepo(db *sql.DB) *TodoRepo {
	return &TodoRepo{db: db}
}

func (r *TodoRepo) Create(todo *Todo) (*Todo, error) {
	query := "INSERT INTO todos (title, completed) VALUES ($1, $2) RETURNING id"
	err := r.db.QueryRow(query, todo.Title, todo.Completed).Scan(&todo.ID)
	if err != nil {
		return nil, err
	}
	return todo, nil
}
