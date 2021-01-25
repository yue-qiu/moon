## moon

moon is a toy web framework written in Golang.

### Usage

```Golang
package main

import (
	"fmt"
	"github.com/yue-qiu/moon"
)

func main() {
	r := moon.Default()

	// router
	r.Add("/test", func(ctx *moon.Context) {
		ctx.Write([]byte("hello world"))
	}, []string{"GET"})

	r.Add("/upload", func(ctx *moon.Context) {
		form := ctx.GetMultipartForm()
		if fileHeader, ok := form.File["file"]; ok {
			fmt.Println(fileHeader.Filename)
                }
	}, []string{"POST"})

	r.Run(":8000")
}
```

`Run()` uses port 8080 by default if no parameters are specified . GET and POST are supported at present.