package mnkiserver

import "net/http"

type MnkiHandler = func(req *http.Request, res http.ResponseWriter, params map[string]string)

type RouteHandler map[string]MnkiHandler

func (ms *MnkiServer) Get(route string, handler func(req *http.Request, res http.ResponseWriter, params map[string]string)) {
	ms.RoutesMapper["GET"][route] = handler
}

func (ms *MnkiServer) Post(route string, handler func(req *http.Request, res http.ResponseWriter, params map[string]string)) {
	ms.RoutesMapper["POST"][route] = handler
}

func (ms *MnkiServer) Put(route string, handler func(req *http.Request, res http.ResponseWriter, params map[string]string)) {
	ms.RoutesMapper["PUT"][route] = handler
}

func (ms *MnkiServer) Patch(route string, handler func(req *http.Request, res http.ResponseWriter, params map[string]string)) {
	ms.RoutesMapper["PATCH"][route] = handler
}

func (ms *MnkiServer) Delete(route string, handler func(req *http.Request, res http.ResponseWriter, params map[string]string)) {
	ms.RoutesMapper["DELETE"][route] = handler
}
