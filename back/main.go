package main

import (
	"net/http"

	"github.com/MonkiG/go-todo-react/back/handlers"
	"github.com/MonkiG/go-todo-react/back/mnkiserver"
	"github.com/MonkiG/go-todo-react/back/models"
	"github.com/MonkiG/go-todo-react/back/response"
)

func main() {
	srv := mnkiserver.New(8080)
	srv.UseCors([]string{"*"}, "GET, POST, PUT, DELETE, OPTIONS", "Content-Type, Authorization")
	db := &models.Db{
		Todos: make([]models.Todo, 0),
	}

	srv.Get("/", func(req *http.Request, res http.ResponseWriter) {
		response.JSON(res, 200, map[string]string{"message": "Hello, world!"})
	})

	todoHandler := &handlers.TodoHandler{
		Db: db,
	}
	srv.Get("/api/v1/todo", todoHandler.GetAll)
	srv.Get("/api/v1/todo/{id}", todoHandler.GetById)
	srv.Post("/api/v1/todo", todoHandler.Create)
	srv.Patch("/api/v1/todo/{id}", todoHandler.Update)
	srv.Delete("/api/v1/todo/{id}", todoHandler.Delete)

	srv.Run()
}
