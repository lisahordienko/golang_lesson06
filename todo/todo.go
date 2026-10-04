// Package todo provides functions to persist a simple todo list to disk
// as JSON.
//
// Homework — Task 1 (Lesson 6: File I/O, JSON and Testing):
// Implement SaveTodos and LoadTodos below so that all tests in
// todo_test.go pass, and reach at least 80% statement coverage for
// this package (checked automatically by CI — see the repository
// README for how to run it locally).
package todo

import (
	"encoding/json"
	"fmt"
	"os"
	"time"
)

// Todo represents a single todo-list item.
type Todo struct {
	ID        int       `json:"id"`
	Title     string    `json:"title"`
	Done      bool      `json:"done"`
	CreatedAt time.Time `json:"created_at"`
}

// SaveTodos writes the given todos to the file at path as JSON,
// creating the file if it does not exist and overwriting it if it does.
//
// Errors from JSON encoding and file writes are returned with context.
func SaveTodos(path string, todos []Todo) error {
	data, err := json.Marshal(todos)
	if err != nil {
		return fmt.Errorf("SaveTodos: marshal todos: %w", err)
	}
	if err := os.WriteFile(path, data, 0o644); err != nil {
		return fmt.Errorf("SaveTodos: write file: %w", err)
	}
	return nil
}

// LoadTodos reads and parses the todo list stored at path.
//
// Read and JSON parsing errors are returned with context while preserving
// their underlying types for inspection with errors.Is and errors.As.
func LoadTodos(path string) ([]Todo, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("LoadTodos: read file: %w", err)
	}

	var todos []Todo
	if err := json.Unmarshal(data, &todos); err != nil {
		return nil, fmt.Errorf("LoadTodos: parse JSON: %w", err)
	}
	return todos, nil
}
