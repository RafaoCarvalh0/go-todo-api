package store

import (
	"go-todo-api/todos/inputs"
	"go-todo-api/todos/models"
)

func GetTodos() []models.Todo {
	todos := make([]models.Todo, 0, len(models.Todos))

	for _, todo := range models.Todos {
		todos = append(todos, todo)
	}

	return todos
}

func GetTodoByID(id int) (models.Todo, bool) {
	todo, exists := models.Todos[id]

	return todo, exists
}

func CreateTodo(createTodoInput inputs.CreateTodoInput) models.Todo {
	lastTodo := getLastTodo()

	newTodoID := lastTodo.ID + 1

	newTodo := models.Todo{
		ID:    newTodoID,
		Title: createTodoInput.Title,
		Done:  createTodoInput.Done,
	}

	models.Todos[newTodoID] = newTodo

	return newTodo
}

func getLastTodo() models.Todo {
	var lastID int
	for id := range models.Todos {
		if id > lastID {
			lastID = id
		}
	}
	return models.Todos[lastID]
}

func DeleteTodo(id int) {
	delete(models.Todos, id)
}

func UpdateTodo(todo models.Todo, updateTodoInput inputs.UpdateTodoInput) models.Todo {
	updatedTodo := getTodoWithChanges(todo, updateTodoInput)

	models.Todos[todo.ID] = updatedTodo

	return updatedTodo
}

func getTodoWithChanges(todo models.Todo, input inputs.UpdateTodoInput) models.Todo {
	updatedTodo := todo

	if input.Title != nil {
		updatedTodo.Title = *input.Title
	}

	if input.Done != nil {
		updatedTodo.Done = *input.Done
	}

	return updatedTodo
}
