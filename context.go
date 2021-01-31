package moon

import (
	"mime/multipart"
	"net/http"
)

const MAXMULTIPARTYMEM = 8 << 20 // 8MB

type Context struct {
	Rsp http.ResponseWriter
	Req *http.Request
	Params
}

func (c *Context) Init(rsp http.ResponseWriter, req *http.Request) {
	c.Req = req
	c.Rsp = rsp
	c.Params = make(map[string]string)
}

func (c *Context) GetParam(key string) (val string) {
	return c.Params[key]
}

func (c *Context) ParamsCount() (count int) {
	return c.Params.Count()
}

func (c *Context) Write(msg []byte) (int, error) {
	i, err := c.Rsp.Write(msg)
	if err != nil {
		return 0, err
	}
	return i, nil
}

func (c *Context) GetHeaderField(key string) string {
	return c.Req.Header.Get(key)
}

func (c *Context) GetFormField(key string) []string {
	(c.Req).ParseForm()
	return (c.Req).Form[key]
}

func (c *Context) GetDefaultFormField(key string, val string) []string {
	res := c.GetFormField(key)
	if len(res) == 0 {
		return []string{val}
	}
	return res
}

func (c *Context) GetPostFormField(key string) []string {
	(c.Req).ParseForm()
	return (c.Req).PostForm[key]
}

func (c *Context) GetDefaultPostFormField(key string, val string) []string {
	res := c.GetPostFormField(key)
	if len(res) == 0 {
		return []string{val}
	}
	return res
}

func (c *Context) GetQueryField(key string) string {
	return c.Req.URL.Query().Get(key)
}

func (c *Context) GetDefaultQueryField(key string, val string) string {
	res := c.GetQueryField(key)
	if res == "" {
		return val
	}
	return res
}

func (c *Context) GetMultipartForm() *multipart.Form {
	(c.Req).ParseMultipartForm(MAXMULTIPARTYMEM)
	return c.Req.MultipartForm
}

func (c *Context) Copy() Context {
	cp := Context{
		Rsp: c.Rsp,
		Req: c.Req,
	}

	return cp
}
