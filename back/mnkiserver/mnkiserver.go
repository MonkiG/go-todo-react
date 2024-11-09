package mnkiserver

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/MonkiG/go-todo-react/back/response"
)

type RoutesMapper map[string]RouteHandler

type MnkiServer struct {
	Port         uint
	Addr         string
	RoutesMapper RoutesMapper
}

func (ms *MnkiServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	method := r.Method
	endpoint := r.URL.Path

	log.Printf(": [%s] \"%s\"", method, endpoint)
	routes, ok := ms.RoutesMapper[method]

	if !ok {
		response.JSON(w, http.StatusMethodNotAllowed, map[string]string{"message": "Method not allowed"})
		return
	}

	for k, handler := range routes {
		params := ms.matchPattern(k, endpoint)
		handler(r, w, params)
		return
	}
}

// Method that validates the params of the routes
func (ms *MnkiServer) matchPattern(pattern string, path string) map[string]string {
	patternSplitted := strings.Split(pattern, "/")
	pathSplitted := strings.Split(path, "/")
	params := make(map[string]string)

	if (len(patternSplitted) != len(pathSplitted)) || (path == "/" && pattern != "/") {
		return params
	}

	for i, part := range patternSplitted {
		if strings.HasPrefix(part, ":") {
			paramName := part[1:]
			params[paramName] = pathSplitted[i]
		}
	}

	return params
}

func New(port uint) *MnkiServer {
	parsedAddr := fmt.Sprintf("localhost:%d", port)
	return &MnkiServer{
		Port: port,
		Addr: parsedAddr,
		RoutesMapper: map[string]RouteHandler{
			"GET":    make(RouteHandler),
			"POST":   make(RouteHandler),
			"PUT":    make(RouteHandler),
			"PATCH":  make(RouteHandler),
			"DELETE": make(RouteHandler),
		},
	}
}
func (ms *MnkiServer) Run() {
	log.Println("Listening in " + ms.Addr)
	log.Fatal(http.ListenAndServe(ms.Addr, ms))

}
