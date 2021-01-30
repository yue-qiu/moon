package moon

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestParseNamedParam(t *testing.T) {
	_, _, ok := ParseNamedParam("/hello")
	assert.False(t, ok)

	name, start, ok := ParseNamedParam("/hello/:name")
	assert.True(t, ok)
	fmt.Println(name, start)

	_, _, ok = ParseNamedParam("/:")
	assert.False(t, ok)
}

func TestNode_AddRouter(t *testing.T) {
	r := &Tree{
		children: make([]*Tree, 0),
	}

	func1 := func(ctx *Context) {}
	func2 := func(ctx *Context) {}
	func3 := func(ctx *Context) {}
	func4 := func(ctx *Context) {}

	r.AddRouter("/hello/:name/:surname", func1)
	r.AddRouter("/:name/bye", func2)
	r.AddRouter("/run/:name", func3)
	r.AddRouter("/run/flask", func4)
	assert.Equal(t, r.Has("/jack/bye"), true)
	//r.AddRouter("/run/:animal", nil)
}