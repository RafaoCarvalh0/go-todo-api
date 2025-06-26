package todos

import (
	"errors"
	"go-todo-api/todos/external"
	"go-todo-api/todos/inputs"
	"go-todo-api/todos/store"
)

func ListTodos() []external.TodoExternal {
	todos := store.GetTodos()

	externalTodos := make([]external.TodoExternal, len(todos))

	for i := range todos {
		externalTodos[i] = external.FromModel(todos[i])
	}

	return externalTodos
}

func CreateTodo(createTodoInput inputs.CreateTodoInput) external.TodoExternal {
	createdTodoModel := store.CreateTodo(createTodoInput)

	return external.FromModel(createdTodoModel)
}

func DeleteTodo(todoID int) error {
	if _, exists := store.GetTodoByID(todoID); exists {
		store.DeleteTodo(todoID)
		return nil
	} else {
		return errors.New("todo not found")
	}
}

func UpdateTodo(todoID int, updateTodoInput inputs.UpdateTodoInput) (external.TodoExternal, error) {
	if existentTodo, exists := store.GetTodoByID(todoID); exists {
		updatedTodo := store.UpdateTodo(existentTodo, updateTodoInput)
		return external.FromModel(updatedTodo), nil
	} else {
		return external.TodoExternal{}, errors.New("todo not found")
	}
}
