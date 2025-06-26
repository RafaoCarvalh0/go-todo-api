package routes

import (
	"go-todo-api/controllers/todo"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {
	todos := r.Group("/todos")
	{
		todos.GET("/", todo.ListTodos)
		todos.POST("/create", todo.CreateTodo)
		todos.DELETE("/:id", todo.DeleteTodo)
		todos.PUT("/:id", todo.UpdateTodo)
	}
}
