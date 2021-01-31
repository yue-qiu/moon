package moon

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"net/http"
	"strings"
	"testing"
)

func TestEngine_Run(t *testing.T) {
	go func() {
		r := Default()

		r.Add("/testGet", func(ctx *Context) {
			ctx.Write([]byte("hello world"))
		}, []string{"GET"})

		r.Add("/testPost", func(ctx *Context) {
			ctx.Write([]byte(ctx.GetDefaultFormField("name", "rose")[0]))
		}, []string{"POST"})

		r.Add("/:name/param", func(ctx *Context) {
			ctx.Write([]byte(ctx.Params.Get("name")))
		}, []string{"GET"})
		r.Run()
	}()

	// test GET
	res, err := http.Get("http://127.0.0.1:8080/testGet")
	if err != nil {
		fmt.Println(err)
	}
	rsp, _ := ioutil.ReadAll(res.Body)
	assert.Equal(t, "hello world", string(rsp))

	// test POST
	res,err = http.Post("http://127.0.0.1:8080/testPost",
		"application/x-www-form-urlencoded",
		strings.NewReader("name=jack"))
	if err != nil {
		fmt.Println(err)
	}
	rsp, _ = ioutil.ReadAll(res.Body)
	assert.Equal(t, "jack", string(rsp))

	// test named params
	res, err = http.Get("http://127.0.0.1:8080/rowan/param")
	if err != nil {
		fmt.Println(err)
	}
	rsp, _ = ioutil.ReadAll(res.Body)
	assert.Equal(t, "rowan", string(rsp))
	assert.NotEqual(t, "kelly", string(rsp))
}
