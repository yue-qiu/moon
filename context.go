package moon

import "net/http"


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

func (c *Context) GetFormKey(key string) []string {
	(c.req).ParseForm()
	return (c.req).Form[key]
}
