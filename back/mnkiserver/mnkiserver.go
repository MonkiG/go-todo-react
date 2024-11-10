package mnkiserver

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/MonkiG/go-todo-react/back/response"
)

type RoutesMapper map[string]RouteHandler

type cors struct {
	allowed bool
	origins []string
	methods string
	headers string
}
type MnkiServer struct {
	Port         uint
	Addr         string
	RoutesMapper RoutesMapper
	cors         *cors
}

func (ms *MnkiServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if ms.cors.allowed {
		for _, o := range ms.cors.origins {
			w.Header().Set("Access-Control-Allow-Origin", o)
		}
		w.Header().Set("Access-Control-Allow-Methods", ms.cors.methods)
		w.Header().Set("Access-Control-Allow-Headers", ms.cors.headers)
	}

	method := r.Method
	endpoint := r.URL.Path

	log.Printf(": [%s] \"%s\"", method, endpoint)
	routes, ok := ms.RoutesMapper[method]

	if !ok {
		response.JSON(w, http.StatusMethodNotAllowed, map[string]string{"message": "Method not allowed"})
		return
	}

	for k, handler := range routes {
		match := ms.matchPattern(r, k, endpoint)
		if match {
			handler(r, w)
			return
		}
	}

	response.JSON(w, http.StatusNotFound, map[string]string{"message": "Route not found"})
}

// Method that validates the params of the routes
func (ms *MnkiServer) matchPattern(r *http.Request, pattern string, path string) bool {
	patternSplitted := strings.Split(pattern, "/")
	pathSplitted := strings.Split(path, "/")
	// params := make(map[string]string)

	if len(patternSplitted) != len(pathSplitted) {
		return false
	}

	for i, part := range patternSplitted {
		if strings.HasPrefix(part, "{") && strings.HasSuffix(part, "}") {
			paramName := part[1 : len(part)-1]
			paramValue := pathSplitted[i]
			r.SetPathValue(paramName, paramValue)
			continue
		} else if part != pathSplitted[i] {
			return false
		}
	}

	return true
}

func (ms *MnkiServer) UseCors(allowedOrigins []string, allowedMethods string, allowedHeaders string) {
	ms.cors.allowed = true
	ms.cors.origins = allowedOrigins
	ms.cors.methods = allowedMethods
	ms.cors.headers = allowedHeaders
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
		cors: &cors{
			allowed: false,
			methods: "",
			origins: nil,
			headers: "",
		},
	}
}
func (ms *MnkiServer) Run() {
	log.Println("Listening in " + ms.Addr)
	log.Fatal(http.ListenAndServe(ms.Addr, ms))

}
