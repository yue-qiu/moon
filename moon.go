package moon

import (
	"net/http"
	"strings"
	"sync"
)

var HTTPMETHODS = map[string]bool{"GET": true, "POST": true, "PUT": true, "HEAD": true, "DELETE": true,
	"OPTIONS": true, "CONNECT": true, "TRACE": true, "PATCH": true}

type Engine struct {
	ft      map[string]*Tree
	pool    sync.Pool
}


func (e *Engine) Add(pattern string, handle Handler, methods MethodList) error {
	for _, method := range methods {
		method = strings.ToUpper(method)
		if HTTPMETHODS[method] == false {
			continue
		}
		if e.ft[method] == nil {
			e.ft[method] = &Tree{
				children: make([]*Tree, 0),
			}
		}
		e.ft[method].AddRouter(pattern, handle)
	}
	return nil
}

func (e *Engine) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	ctx := e.pool.Get().(*Context)
	ctx.Init(w, r)

	handle := e.ft[r.Method].Retrieve(ctx.Req.URL.Path, ctx)
	handle(ctx)

	e.pool.Put(ctx)
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
	engine := &Engine{
		ft: make(map[string]*Tree),
	}
	engine.pool.New = func() interface{} {
		return new(Context)
	}

	return engine
}
