package main

import (
	"net/http"

	"github.com/MonkiG/go-todo-react/back/handlers"
	"github.com/MonkiG/go-todo-react/back/mnkiserver"
	"github.com/MonkiG/go-todo-react/back/response"
)

func main() {
	srv := mnkiserver.New(8080)
	db := make([]string, 0)

	srv.Get("/", func(req *http.Request, res http.ResponseWriter) {
		response.JSON(res, 200, map[string]string{"message": "Hello, world!"})
	})

	todoHandler := &handlers.TodoHandler{
		Db: db,
	}

	srv.Get("/todo", todoHandler.GetAll)
	srv.Run()
}
