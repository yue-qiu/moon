package moon

import (
	"mime/multipart"
	"net/http"
)

const MAXMULTIPARTYMEM = 8 << 20 // 8MB

type Context struct {
	rsp http.ResponseWriter
	req *http.Request
}

func (c *Context) Write(msg []byte) (int, error) {
	i, err := c.rsp.Write(msg)
	if err != nil {
		return 0, err
	}
	return i, nil
}

func (c *Context) Method() string {
	return c.req.Method
}

func (c *Context) GetHeaderField(key string) string {
	return c.req.Header.Get(key)
}

func (c *Context) GetFormField(key string) []string {
	(c.req).ParseForm()
	return (c.req).Form[key]
}

func (c *Context) GetDefaultFormField(key string, val string) []string {
	res := c.GetFormField(key)
	if len(res) == 0 {
		return []string{val}
	}
	return res
}

func (c *Context) GetPostFormField(key string) []string {
	(c.req).ParseForm()
	return (c.req).PostForm[key]
}

func (c *Context) GetDefaultPostFormField(key string, val string) []string {
	res := c.GetPostFormField(key)
	if len(res) == 0 {
		return []string{val}
	}
	return res
}

func (c *Context) GetQueryField(key string) string {
	return c.req.URL.Query().Get(key)
}

func (c *Context) GetDefaultQueryField(key string, val string) string {
	res := c.GetQueryField(key)
	if res == "" {
		return val
	}
	return res
}

func (c *Context) GetMultipartForm() *multipart.Form {
	(c.req).ParseMultipartForm(MAXMULTIPARTYMEM)
	return c.req.MultipartForm
}

func (c *Context) GetProto() string {
	return c.req.Proto
}
