## moon

moon is a toy web framework written in Golang

### Usage

```Golang
package main

import "github.com/yue-qiu/moon"

func main() {
	r := moon.Default()
	
	// router
	r.Add("/test", func(ctx *moon.Context) {
            ctx.Write([]byte("hello world")) 
	}, []string{"GET"})
	
	r.Run()
}
```

Port 8080 is used by default. GET is the only supported method at present.