## moon

moon is a simple web framework, stores route paths with radix tree, supports named parameter such as `/hello/:name`.

Trailing slash will be add to the end of path if it is missing. eg: `/hello/world` will be modifying to `/hello/world/`.

### Usage

```Golang
package main

import (
	"fmt"
	"github.com/yue-qiu/moon"
)

func main() {
	r := moon.Default()

	// add path
	r.Add("/test", func(ctx *moon.Context) {
		ctx.Write([]byte("hello world"))
	}, []string{"GET"})

	r.Add("/upload", func(ctx *moon.Context) {
		form := ctx.GetMultipartForm()
		if fileHeader, ok := form.File["file"]; ok {
			fmt.Println(fileHeader.Filename)
                }
	}, []string{"POST"})
	
	r.Add("/:name/param", func(ctx *Context) {
		ctx.Write([]byte(ctx.Params.Get("name")))
	}, []string{"GET"})
	
	r.Run(":8000")
}
```

`Run()` uses port 8080 by default if no parameters are specified. All HTTP methods are supported.