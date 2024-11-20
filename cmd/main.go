package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/reon150/go-todo-rest-api/config"
	"github.com/reon150/go-todo-rest-api/internal/tasks"
	"github.com/reon150/go-todo-rest-api/internal/todos"
)

func taskPathHandler(todoHandler *todos.TodoHandler, taskHandler *tasks.TaskHandler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		segments := strings.Split(strings.Trim(r.URL.Path, "/"), "/")

		switch {
		case len(segments) == 1 && segments[0] == "todo":
			todoHandler.ServeHTTP(w, r)
		case len(segments) == 2 && segments[0] == "todo":
			todoHandler.ServeHTTP(w, r)
		case len(segments) == 3 && segments[0] == "todo" && segments[2] == "task":
			taskHandler.ServeHTTP(w, r)
		case len(segments) == 4 && segments[0] == "todo" && segments[2] == "task":
			taskHandler.ServeHTTP(w, r)
		default:
			http.NotFound(w, r)
		}
	})
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/" {
		http.Redirect(w, r, "/todo", http.StatusPermanentRedirect)
		return
	}
	http.NotFound(w, r)
}

func main() {
	config.LoadConfig()
	config.InitDatabase()

	// Initialize Todo components
	todoRepo := todos.NewTodoRepository(config.DB)
	todoService := todos.NewTodoService(todoRepo)
	todoHandler := todos.NewTodoHandler(todoService)

	// Initialize Task components
	taskRepo := tasks.NewTaskRepository(config.DB)
	taskService := tasks.NewTaskService(taskRepo)
	taskHandler := tasks.NewTaskHandler(taskService)

	// Set up routing
	mainHandler := taskPathHandler(todoHandler, taskHandler)

	// Register handlers
	http.HandleFunc("/", rootHandler)
	http.Handle("/todo", mainHandler)
	http.Handle("/todo/", mainHandler)

	fmt.Printf("Starting server on port 8080...\n")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
