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

func (r *TodoRepo) GetAll() ([]*Todo, error) {
	query := "SELECT id, title, completed FROM todos"
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var todos []*Todo
	for rows.Next() {
		todo := &Todo{}
		err := rows.Scan(&todo.ID, &todo.Title, &todo.Completed)
		if err != nil {
			return nil, err
		}
		todos = append(todos, todo)
	}
	return todos, nil
}

func (r *TodoRepo) GetByID(id int) (*Todo, error) {
	query := "SELECT id, title, completed FROM todos WHERE id = $1"
	todo := &Todo{}
	err := r.db.QueryRow(query, id).Scan(&todo.ID, &todo.Title, &todo.Completed)
	if err != nil {
		return nil, err
	}
	return todo, nil
}

func (r *TodoRepo) Update(todo *Todo) error {
	query := "UPDATE todos SET title = $1, completed = $2 WHERE id = $3"
	_, err := r.db.Exec(query, todo.Title, todo.Completed, todo.ID)
	return err
}

func (r *TodoRepo) Delete(id int) error {
	query := "DELETE FROM todos WHERE id = $1"
	_, err := r.db.Exec(query, id)
	return err
}
