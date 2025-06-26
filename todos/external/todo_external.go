package external

import (
	"go-todo-api/todos/models"
)

type TodoExternal struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
	Done  bool   `json:"done"`
}

func FromModel(todo models.Todo) TodoExternal {
	return TodoExternal{
		ID:    todo.ID,
		Title: todo.Title,
		Done:  todo.Done,
	}
}
