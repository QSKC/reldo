package main

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

type Todo struct {
	ID   int
	Task string
	Done bool
}

const dataFile = "todos.json"

// load todos from file
func loadTodos() ([]Todo, error) {
	var todos []Todo
	file, err := ioutil.ReadFile(dataFile)
	if err != nil {
		if os.IsNotExist(err) {
			return []Todo{}, nil
		}
		return nil, err
	}
	if len(file) > 0 {
		if err := json.Unmarshal(file, &todos); err != nil {
			return nil, err
		}
	}
	return todos, nil
}

// Save todos to file
func saveTodos(todos []Todo) error {
	data, err := json.MarshalIndent(todos, "", " ")
	if err != nil {
		return err
	}
	return ioutil.WriteFile(dataFile, data, 0644)
}

// Add new task

func addTask(task string) error {
	todos, err := loadTodos()
	if err != nil {
		return err
	}
	id := len(todos) + 1
	todos = append(todos, Todo{ID: id, Task: task, Done: false})
	return saveTodos(todos)
}

func listTasks() ([]Todo, error) {
	return loadTodos()
}

// Mark a task done
func markDone(id int) error {
	todos, err := loadTodos()
	if err != nil {
		return err
	}
	for i := range todos {
		if todos[i].ID == id {
			todos[i].Done = true
		}
	}
	return saveTodos(todos)
}

// EDIT A TASK, PLZ
func editTask(toid int, task string) error {
	todos, err := loadTodos()
	if err != nil {
		return err
	}
	for i := range todos {
		if todos[i].ID == toid {
			todos[i].Task = task
		}
	}
	return saveTodos(todos)
}

// Remove a task
func removeTask(id int) error {
	todos, err := loadTodos()
	if err != nil {
		return err
	}
	newTodos := []Todo{}
	for _, t := range todos {
		if t.ID != id {
			newTodos = append(newTodos, t)
		}
	}
	// Reassign IDs for consistency
	for i := range newTodos {
		newTodos[i].ID = i + 1
	}
	return saveTodos(newTodos)
}
