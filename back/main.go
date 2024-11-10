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
	srv.Get("/todo", todoHandler.GetAll)
	srv.Get("/todo/{id}", todoHandler.GetById)
	srv.Post("/todo", todoHandler.Create)
	srv.Patch("/todo/{id}", todoHandler.Update)
	srv.Delete("/todo/{id}", todoHandler.Delete)

	srv.Run()
}
