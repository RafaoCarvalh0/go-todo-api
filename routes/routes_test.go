package routes

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestTodosGroupExists(t *testing.T) {
	gin.SetMode(gin.TestMode)
	r := gin.New()

	SetupRoutes(r)

	w := httptest.NewRecorder()

	req, _ := http.NewRequest("GET", "/todos/", nil)
	r.ServeHTTP(w, req)

	if w.Code == http.StatusNotFound {
		t.Error("Expected route /todos/ to be registered, but got 404")
	}
}

func TestAllRoutesRegistered(t *testing.T) {
	gin.SetMode(gin.TestMode)
	r := gin.New()

	SetupRoutes(r)

	testCases := []struct {
		method string
		path   string
		name   string
	}{
		{"GET", "/todos/", "ListTodos"},
		{"POST", "/todos/create", "CreateTodo"},
		{"DELETE", "/todos/1", "DeleteTodo"},
		{"PUT", "/todos/1", "UpdateTodo"},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			w := httptest.NewRecorder()
			req, _ := http.NewRequest(tc.method, tc.path, nil)
			r.ServeHTTP(w, req)

			if w.Code == http.StatusNotFound {
				t.Errorf("Expected route %s %s to be registered, but got 404", tc.method, tc.path)
			}
		})
	}
}

func TestRoutesWithInvalidPaths(t *testing.T) {
	gin.SetMode(gin.TestMode)
	r := gin.New()

	SetupRoutes(r)

	testCases := []struct {
		method string
		path   string
		name   string
	}{
		{"GET", "/invalid", "Invalid GET route"},
		{"POST", "/todos", "Invalid POST route (missing /create)"},
		{"DELETE", "/todos", "Invalid DELETE route (missing ID)"},
		{"PUT", "/todos", "Invalid PUT route (missing ID)"},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			w := httptest.NewRecorder()
			req, _ := http.NewRequest(tc.method, tc.path, nil)
			r.ServeHTTP(w, req)

			if w.Code != http.StatusNotFound {
				t.Errorf("Expected route %s %s to return 404, but got %d", tc.method, tc.path, w.Code)
			}
		})
	}
}
