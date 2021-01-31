package moon

import (
	"fmt"
)

type NodeType uint8

const (
	STATIC NodeType = iota
	PARAM
)

type Tree struct {
	path string
	nType NodeType
	indices string
	children []*Tree
	handle Handler
}

func longestCommonPrefix(a, b string) (idx int) {
	minLen := len(a)
	if len(b) < minLen {
		minLen = len(b)
	}

	for idx < minLen && a[idx] == b[idx] {
		idx++
	}

	return idx
}

func (n *Tree) getIndex(c byte) int {
	for i := range n.indices {
		if n.indices[i] == c {
			return i
		}
	}
	return -1
}

// 完成路由路径与处理器的绑定
// 统一使用 [proto]://[domain]/path/ 的形式
func (n *Tree) AddRouter(path string, handler Handler) {
	if path == "" {
		panic("empty path")
	}
	if path[0] != '/' {
		panic("path have to start with /")
	}
	if path[len(path)-1] != '/' {
		path += "/"
	}
	n.insert(path, handler)
}

func (n *Tree) insert(path string, handle Handler) {
	if len(path) == 0 {
		return
	}
	i := longestCommonPrefix(n.path, path)
	if n.path == path || (i > 0 && path[i-1] == ':') {
		panic(fmt.Sprintf("%s conflict with %s\n", n.path, path))
	}

	// split the original Node
	if i < len(n.path) {
		child := &Tree{
			handle: n.handle,
			indices: n.indices,
			path: n.path[i:],
			children: make([]*Tree, 0),
		}
		for idx := range n.indices {
			child.children = append(child.children, n.children[idx])
		}
		n.indices = string(child.path[0])
		n.path = n.path[: i]
		n.children = []*Tree{child}
		n.handle = nil
	}

	// generate a new Node
	if i < len(path) {
		idx := n.getIndex(path[i])
		if idx != -1 {
			n.children[idx].insert(path[i:], handle)
		} else {
			child := &Tree{
				path:     path[i:],
				handle:   handle,
				children: make([]*Tree, 0),
			}
			child.divideNodeByParam()
			n.indices += string(child.path[0])
			n.children = append(n.children, child)
		}
	}
}

func (n *Tree) Retrieve(path string, ctx *Context) Handler {
	if len(path) > 0 && path[len(path)-1] != '/' {
		path += "/"
	}

	i := longestCommonPrefix(n.path, path)
	if i == len(path) && n.nType == STATIC {
		return n.handle
	}

	// 匹配命名参数与真实路径
	if n.nType == PARAM {
		p := 0
		for p < len(path) && path[p] != '/' {
			p++
		}
		ctx.Params[n.path[1:]] = path[: p]
		path = path[p:]
		idx := n.getIndex(path[0])
		return n.children[idx].Retrieve(path, ctx)
	}

	idx := n.getIndex(path[i])
	if idx == -1 {
		if paramIdx := n.getIndex(':'); paramIdx != -1 {
			return n.children[paramIdx].Retrieve(path[i:], ctx)
		}
		return nil
	}
	return n.children[idx].Retrieve(path[i:], ctx)
}

// 根据节点路径是否含命名参数更新状态
func (n *Tree) divideNodeByParam() {
	param, start, ok := ParseNamedParam(n.path)
	if param == n.path {
		return
	}
	if !ok {
		n.nType = STATIC
		return
	}
	if start == 0 {
		child := &Tree{
			path: n.path[len(param):],
			children: make([]*Tree, 0),
			handle: n.handle,
		}
		child.divideNodeByParam()
		n.path = n.path[: len(param)]
		n.handle = nil
		n.indices += string(child.path[0])
		n.nType = PARAM
		n.children = append(n.children, child)
	} else {
		child := &Tree{
			path:     n.path[start:],
			nType:    PARAM,
			children: make([]*Tree, 0),
			handle:   n.handle,
		}
		child.divideNodeByParam()
		n.path = n.path[: start]
		n.handle = nil
		n.indices += string(child.path[0])
		n.nType = STATIC
		n.children = append(n.children, child)
	}
}

// 解析 path 中第一个命名参数及其位置
func ParseNamedParam(path string) (name string, start int, ok bool) {
	if len(path) == 0 {
		return
	}
	start = 0
	for start < len(path) && path[start] != ':' {
		start++
	}


	end := start
	for end < len(path) && path[end] != '/' {
		end++
	}
	return path[start: end], start, 1 < (end - start)
}

