package todo

import (
	"net/http"
	"strconv"

	"go-todo-api/todos/inputs"

	"go-todo-api/todos"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func ListTodos(c *gin.Context) {
	c.JSON(http.StatusOK, todos.ListTodos())
}

func CreateTodo(c *gin.Context) {
	var input inputs.CreateTodoInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON"})
		return
	}

	validate := validator.New()
	if err := validate.Struct(input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid data", "details": err.Error()})
		return
	}

	newTodo := todos.CreateTodo(input)

	c.JSON(http.StatusCreated, newTodo)
}

func DeleteTodo(c *gin.Context) {
	id := c.Param("id")
	todoID, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	if err := todos.DeleteTodo(todoID); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Todo deleted successfully"})
}

func UpdateTodo(c *gin.Context) {
	id := c.Param("id")
	todoID, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	var input inputs.UpdateTodoInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON"})
		return
	}

	validate := validator.New()
	if err := validate.Struct(input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid data", "details": err.Error()})
		return
	}

	updatedTodo, err := todos.UpdateTodo(todoID, input)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, updatedTodo)
}
