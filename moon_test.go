package moon

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"net/http"
	"testing"
)

func TestEngine_Run(t *testing.T) {
	go func() {
		r := Default()
		r.Add("/test", func(ctx *Context) {
			ctx.Write([]byte("hello world"))
		}, []string{"GET"})

		r.Run()
	}()

	res, err := http.Get("http://127.0.0.1:8080/test")
	if err != nil {
		fmt.Println(err)
	}
	rsp, _ := ioutil.ReadAll(res.Body)
	assert.Equal(t, "hello world", string(rsp))
}
