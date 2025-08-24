package main

import (
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAPI_HealthCheck(t *testing.T) {
	server := NewServer()
	rec := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "/health", nil)
	server.ServeHTTP(rec, req)

	assert.Equal(t, 200, rec.Code)
	assert.Equal(t, "OK", rec.Body.String())
}

func TestAPI_CreateTodo(t *testing.T) {
	server := NewServer()
	rec := httptest.NewRecorder()
	req := httptest.NewRequest("POST", "/todos", nil)
	server.ServeHTTP(rec, req)

	assert.Equal(t, 201, rec.Code)
	assert.Equal(t, `{"id":1,"title":"Sample Todo","completed":false}`, rec.Body.String())
}
