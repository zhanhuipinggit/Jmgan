package gan

import (
	"fmt"
	"net/http"
)

type HandlerFunc func(http.ResponseWriter, *http.Request)

type Engine struct {
	router map[string]HandlerFunc
}

func New() *Engine {
	return &Engine{router: make(map[string]HandlerFunc)}
}
func (engine *Engine) Route(method string, pattern string, handler HandlerFunc) {
	key := method + "-" + pattern
	engine.router[key] = handler
}

func (engine *Engine) GET(pattern string, handlerFunc HandlerFunc) {
	engine.Route("GET", pattern, handlerFunc)
}

func (engine *Engine) POST(patter string, handlerFunc HandlerFunc) {
	engine.Route("POST", patter, handlerFunc)
}

func (engine *Engine) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	key := req.Method + "-" + req.URL.Path
	if handler, ok := engine.router[key]; ok {
		handler(w, req)
	} else {
		fmt.Fprintf(w, "404 NOT FOUND: %s\n", req.URL)
	}
}

func (engine *Engine) Run(addr string) (error error) {
	return http.ListenAndServe(addr, engine)
}
