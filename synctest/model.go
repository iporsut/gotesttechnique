package main

import "time"

type Todo struct {
	ID          int
	Title       string
	Completed   bool
	CreatedAt   time.Time
	CompletedAt *time.Time
}

func (t *Todo) MarkComplete() {
	now := time.Now()
	t.Completed = true
	t.CompletedAt = &now
}
