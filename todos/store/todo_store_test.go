package store

import (
	"testing"

	"go-todo-api/todos/inputs"
	"go-todo-api/todos/models"
)

func setupTestDB(t *testing.T) {
	originalTodos := models.Todos

	t.Cleanup(func() {
		models.Todos = originalTodos
	})

	models.Todos = map[int]models.Todo{
		1: {ID: 1, Title: "Aprender Go", Done: false},
		2: {ID: 2, Title: "Criar API com Gin", Done: true},
	}
}

func TestGetTodos(t *testing.T) {
	setupTestDB(t)

	result := GetTodos()

	expectedCount := len(models.Todos)
	if len(result) != expectedCount {
		t.Errorf("Expected %d todos, got %d", expectedCount, len(result))
	}

	for _, todo := range result {
		if todo.ID <= 0 {
			t.Errorf("Expected valid ID, got %d", todo.ID)
		}
		if todo.Title == "" {
			t.Errorf("Expected non-empty title, got empty string")
		}
	}
}

func TestGetTodoByID(t *testing.T) {
	setupTestDB(t)

	todoID := 1

	todo, exists := GetTodoByID(todoID)

	if !exists {
		t.Errorf("Expected todo to exist, got false")
	}

	if todo.ID != todoID {
		t.Errorf("Expected ID %d, got %d", todoID, todo.ID)
	}
}

func TestGetTodoByIDNotFound(t *testing.T) {
	setupTestDB(t)

	todoID := 999

	todo, exists := GetTodoByID(todoID)

	if exists {
		t.Errorf("Expected todo to not exist, got true")
	}

	zeroTodo := models.Todo{}
	if todo != zeroTodo {
		t.Errorf("Expected zero Todo, got %+v", todo)
	}
}

func TestCreateTodo(t *testing.T) {
	setupTestDB(t)

	input := inputs.CreateTodoInput{
		Title: "Test Todo",
		Done:  true,
	}

	result := CreateTodo(input)

	if result.Title != input.Title {
		t.Errorf("Expected title '%s', got '%s'", input.Title, result.Title)
	}

	if result.Done != input.Done {
		t.Errorf("Expected done %t, got %t", input.Done, result.Done)
	}

	if result.ID != 3 {
		t.Errorf("Expected ID 3, got %d", result.ID)
	}

	if _, exists := models.Todos[result.ID]; !exists {
		t.Errorf("Expected todo to be added to map, but it doesn't exist")
	}
}

func TestDeleteTodo(t *testing.T) {
	setupTestDB(t)

	input := inputs.CreateTodoInput{
		Title: "Todo to delete",
		Done:  false,
	}
	todo := CreateTodo(input)

	if _, exists := models.Todos[todo.ID]; !exists {
		t.Fatalf("Todo should exist before deletion")
	}

	DeleteTodo(todo.ID)

	if _, exists := models.Todos[todo.ID]; exists {
		t.Errorf("Expected todo to be deleted, but it still exists")
	}
}

func TestUpdateTodo(t *testing.T) {
	setupTestDB(t)

	input := inputs.CreateTodoInput{
		Title: "Original Title",
		Done:  false,
	}
	originalTodo := CreateTodo(input)

	newTitle := "Updated Title"
	newDone := true
	updateInput := inputs.UpdateTodoInput{
		Title: &newTitle,
		Done:  &newDone,
	}

	result := UpdateTodo(originalTodo, updateInput)

	if result.Title != newTitle {
		t.Errorf("Expected title '%s', got '%s'", newTitle, result.Title)
	}

	if result.Done != newDone {
		t.Errorf("Expected done %t, got %t", newDone, result.Done)
	}

	if result.ID != originalTodo.ID {
		t.Errorf("Expected ID %d, got %d", originalTodo.ID, result.ID)
	}

	if storedTodo, exists := models.Todos[originalTodo.ID]; !exists {
		t.Errorf("Expected todo to exist in map after update")
	} else if storedTodo.Title != newTitle {
		t.Errorf("Expected stored todo to have updated title '%s', got '%s'", newTitle, storedTodo.Title)
	}
}

func TestApplyChanges(t *testing.T) {
	setupTestDB(t)

	originalTodo := models.Todo{
		ID:    1,
		Title: "Original Title",
		Done:  false,
	}

	newTitle := "New Title"
	newDone := true
	input := inputs.UpdateTodoInput{
		Title: &newTitle,
		Done:  &newDone,
	}

	result := getTodoWithChanges(originalTodo, input)

	if result.Title != newTitle {
		t.Errorf("Expected title '%s', got '%s'", newTitle, result.Title)
	}

	if result.Done != newDone {
		t.Errorf("Expected done %t, got %t", newDone, result.Done)
	}

	if result.ID != originalTodo.ID {
		t.Errorf("Expected ID %d, got %d", originalTodo.ID, result.ID)
	}
}

func TestApplyChangesPartial(t *testing.T) {
	setupTestDB(t)

	originalTodo := models.Todo{
		ID:    1,
		Title: "Original Title",
		Done:  false,
	}

	newTitle := "New Title"
	input := inputs.UpdateTodoInput{
		Title: &newTitle,
	}

	result := getTodoWithChanges(originalTodo, input)

	if result.Title != newTitle {
		t.Errorf("Expected title '%s', got '%s'", newTitle, result.Title)
	}

	if result.Done != originalTodo.Done {
		t.Errorf("Expected done to remain %t, got %t", originalTodo.Done, result.Done)
	}

	if result.ID != originalTodo.ID {
		t.Errorf("Expected ID %d, got %d", originalTodo.ID, result.ID)
	}
}

func TestGetLastTodo(t *testing.T) {
	setupTestDB(t)

	result := getLastTodo()

	if result.ID <= 0 {
		t.Errorf("Expected valid ID, got %d", result.ID)
	}

	maxID := 0
	for id := range models.Todos {
		if id > maxID {
			maxID = id
		}
	}

	if result.ID != maxID {
		t.Errorf("Expected ID %d (max), got %d", maxID, result.ID)
	}
}
