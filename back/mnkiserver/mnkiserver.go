package mnkiserver

import (
	"fmt"
	"log"
	"net/http"

	"github.com/MonkiG/go-todo-react/back/response"
)

type RoutesMapper map[string]RouteHandler

type MnkiServer struct {
	Port         uint
	Addr         string
	RoutesMapper RoutesMapper
}

func (sg *MnkiServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	method := r.Method
	endpoint := r.URL.Path

	log.Printf(": [%s] \"%s\"", method, endpoint)
	_, ok := sg.RoutesMapper[method]

	if !ok {
		response.JSON(w, http.StatusMethodNotAllowed, map[string]string{"message": "Method not allowed"})
		return
	}

	_, ok = sg.RoutesMapper[method][endpoint]

	if !ok {
		response.JSON(w, http.StatusNotFound, map[string]string{"message": "Route not found"})
		return
	}

	sg.RoutesMapper[method][endpoint](r, w)
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
			"DELETE": make(RouteHandler),
		},
	}
}
func (sg *MnkiServer) Run() {
	log.Println("Listening in " + sg.Addr)
	log.Fatal(http.ListenAndServe(sg.Addr, sg))

}
