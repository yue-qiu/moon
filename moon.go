package moon

import (
	"errors"
	"fmt"
	"net/http"
	"strings"
)

const (
	GET = "GET"
	POST = "POST"
)

type Engine struct {
	ftOfGet ForwardingTable
	ftOFPost ForwardingTable
}

func (e *Engine) Add(pattern string, handler Handler, methods MethodList) error {
	for _, method := range methods {
		switch strings.ToUpper(method) {
		case GET:
			e.ftOfGet[pattern] = handler
		case POST:
			e.ftOFPost[pattern] = handler
		default:
			return errors.New(fmt.Sprintf("Unsupported Method: %s\n", method))
		}
	}
	return nil
}

func (e *Engine) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	c := Context{
		Rsp:  w,
		Req:  r,

	}
	switch r.Method {
	case GET:
		if handler, ok := e.ftOfGet[r.URL.Path]; ok {
			handler(&c)
		}
	case POST:
		if handler, ok := e.ftOFPost[r.URL.Path]; ok {
			handler(&c)
		}
	}
}

func (e *Engine) Run(addr ...string) {
	switch len(addr) {
	case 0:
		http.ListenAndServe(":8080", e)
	case 1:
		http.ListenAndServe(addr[0], e)
	default:
		panic("too many parameters")
	}
}

func Default() Router {
	return &Engine{
		ftOfGet: map[string]Handler{},
		ftOFPost: map[string]Handler{},
	}
}
