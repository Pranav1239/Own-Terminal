// internal/todo/todo.go
package todo

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"
)

type Todo struct {
	Tasks []string `json:"tasks"`
}

const todoFile = "todos.json"

// AddTodo adds a new todo item to the list
func AddTodo(task string) error {
	todos, err := loadTodos()
	if err != nil {
		return err
	}

	todos.Tasks = append(todos.Tasks, task)
	return saveTodos(todos)
}

// ViewTodos returns all todo items
func ViewTodos() ([]string, error) {
	todos, err := loadTodos()
	if err != nil {
		return nil, err
	}

	return todos.Tasks, nil
}

// DeleteTodo removes a todo item by its ID (index + 1)
func DeleteTodo(taskID string) error {
	id, err := strconv.Atoi(taskID)
	if err != nil {
		return fmt.Errorf("invalid task ID: %v", err)
	}

	todos, err := loadTodos()
	if err != nil {
		return err
	}

	if id < 1 || id > len(todos.Tasks) {
		return fmt.Errorf("task ID out of range")
	}

	// Remove the task (remember, slice is 0-indexed)
	todos.Tasks = append(todos.Tasks[:id-1], todos.Tasks[id:]...)
	return saveTodos(todos)
}

// loadTodos reads todos from the JSON file
func loadTodos() (*Todo, error) {
	todos := &Todo{}

	file, err := os.ReadFile(todoFile)
	if err != nil {
		if os.IsNotExist(err) {
			return todos, nil // Return empty todos if file doesn't exist
		}
		return nil, err
	}

	err = json.Unmarshal(file, todos)
	if err != nil {
		return nil, err
	}

	return todos, nil
}

// saveTodos writes todos to the JSON file
func saveTodos(todos *Todo) error {
	file, err := json.Marshal(todos)
	if err != nil {
		return err
	}

	return os.WriteFile(todoFile, file, 0644)
}
