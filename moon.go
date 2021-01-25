package moon

import (
	"errors"
	"fmt"
	"net/http"
	"strings"
)

const (
	GET = "GET"
)

type Engine struct {
	ftOfGet ForwardingTable
}

func (e *Engine) Add(pattern string, handler Handler, methods MethodList) error {
	for _, method := range methods {
		switch strings.ToUpper(method) {
		case GET:
			e.ftOfGet[pattern] = handler
		default:
			return errors.New(fmt.Sprintf("Unsupported Method: %s\n", method))
		}
	}
	return nil
}

func (e *Engine) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	c := Context{
		rsp:  w,
		req:  r,
	}
	switch r.Method {
	case GET:
		if handler, ok := e.ftOfGet[r.URL.Path]; ok {
			handler(&c)
		}
	}
}

func (e *Engine) Run() {
	http.ListenAndServe(":8080", e)
}

func Default() Router {
	return &Engine{ftOfGet: make(map[string]Handler)}
}
