package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"github.com/MonkiG/go-todo-react/back/models"
	"github.com/MonkiG/go-todo-react/back/response"
	"github.com/MonkiG/go-todo-react/back/types"
)

type TodoHandler struct {
	Db *models.Db
}

func (th *TodoHandler) GetAll(req *http.Request, res http.ResponseWriter) {
	todos := &th.Db.Todos
	response.JSON(res, 200, todos)
}

func (th *TodoHandler) GetById(req *http.Request, res http.ResponseWriter) {
	id := req.PathValue("id")
	if id == "" {
		response.JSON(res, 404, map[string]string{"message": "You should provide an id"})
		return
	}

	var todo *models.Todo

	for i, t := range th.Db.Todos {
		if t.Id == id {
			todo = &th.Db.Todos[i]
		}
	}

	response.JSON(res, 200, todo)
}

func (th *TodoHandler) Create(req *http.Request, res http.ResponseWriter) {
	var createTodoDto models.CreateTodoDto
	decoder := json.NewDecoder(req.Body)
	err := decoder.Decode(&createTodoDto)

	if err != nil {
		response.JSON(res, 500, map[string]string{"message": "Error decoding the json body data", "error": err.Error()})
		return
	}

	now := time.Now()
	newTodo := &models.Todo{
		Id:        strconv.FormatInt(now.Unix(), 10),
		Title:     createTodoDto.Title,
		Data:      createTodoDto.Data,
		Status:    types.TODO,
		CreatedAt: now,
		UpdatedAt: now,
	}

	th.Db.Todos = append(th.Db.Todos, *newTodo)
	response.JSON(res, 200, newTodo)
}

func (th *TodoHandler) Update(req *http.Request, res http.ResponseWriter) {
	id := req.PathValue("id")
	if id == "" {
		response.JSON(res, 404, map[string]string{"message": "You should provide an id"})
		return
	}

	var todoToUpdate *models.Todo
	for i, t := range th.Db.Todos {
		if t.Id == id {
			todoToUpdate = &th.Db.Todos[i]
			break
		}
	}

	if todoToUpdate == nil {
		response.JSON(res, 404, map[string]string{"message": "Todo not found"})
		return
	}

	decoder := json.NewDecoder(req.Body)
	err := decoder.Decode(&todoToUpdate)

	if err != nil {
		response.JSON(res, 500, map[string]string{"message": "Error decoding the json body data", "error": err.Error()})
		return
	}

	todoToUpdate.UpdatedAt = time.Now()
	response.JSON(res, 200, todoToUpdate)
}

func (th *TodoHandler) Delete(req *http.Request, res http.ResponseWriter) {

	id := req.PathValue("id")
	if id == "" {
		response.JSON(res, 404, map[string]string{"message": "You should provide an id"})
		return
	}

	newTodos := make([]models.Todo, 0, len(th.Db.Todos))
	for _, t := range th.Db.Todos {
		if t.Id != id {
			newTodos = append(newTodos, t)
		}
	}

	th.Db.Todos = newTodos

	response.NoContent(res)
}
