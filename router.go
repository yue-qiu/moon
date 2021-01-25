package moon

type Handler func(ctx *Context)

type ForwardingTable map[string]Handler

type MethodList []string

type Router interface {
	Add(string, Handler, MethodList) error
	Run()
}
