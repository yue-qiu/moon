package moon

type Handler func(ctx *Context)

type MethodList []string

type Router interface {
	Add(string, Handler, MethodList) error
	Run(...string)
}
