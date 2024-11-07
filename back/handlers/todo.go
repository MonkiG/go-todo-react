package handlers

import (
	"net/http"

	"github.com/MonkiG/go-todo-react/back/response"
)

type TodoHandler struct {
	Db []string
}

func (th *TodoHandler) GetAll(req *http.Request, res http.ResponseWriter) {
	response.JSON(res, 200, map[string]string{"message": "all todos route"})
}
