package models

import (
	"time"

	"github.com/MonkiG/go-todo-react/back/types"
)

type Todo struct {
	Id        string           `json:"id"`
	Title     string           `json:"title"`
	Data      string           `json:"data"`
	Status    types.TodoStatus `json:"status"`
	CreatedAt time.Time        `json:"createdAt"`
	UpdatedAt time.Time        `json:"updatedAt"`
}

type CreateTodoDto struct {
	Title string `json:"title"`
	Data  string `json:"data"`
}

type UpdateTodoDto struct {
	Title  *string           `json:"title"`
	Data   *string           `json:"data"`
	Status *types.TodoStatus `json:"status"`
}
