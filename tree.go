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
		child.checkStatus()
		for idx := range n.indices {
			child.children = append(child.children, n.children[idx])
		}
		n.indices = string(n.path[i])
		n.path = n.path[: i]
		n.children = []*Tree{child}
		n.handle = nil
		n.checkStatus()
	}

	// generate a new Node
	if i < len(path) {
		idx := n.getIndex(path[i])
		if idx != -1 {
			n.children[i].insert(path[i:], handle)
		} else {
			child := &Tree{
				path:     path[i:],
				handle:   handle,
				children: make([]*Tree, 0),
			}
			child.checkStatus()
			n.indices += string(path[i])
			n.children = append(n.children, child)
		}
	}
}

func (n *Tree) Has(path string) bool {
	i := longestCommonPrefix(n.path, path)
	if i == len(path) {
		return true
	}
	idx := n.getIndex(path[i])
	if idx == -1 {
		return false
	}
	return n.children[idx].Has(path[i:])
}

func (n *Tree) Retrieve(path string) Handler {
	i := longestCommonPrefix(n.path, path)
	if i == len(path) {
		return n.handle
	}
	idx := n.getIndex(path[i])
	if idx == -1 {
		return nil
	}
	return n.children[idx].Retrieve(path[i:])
}

// 根据节点路径是否含命名参数更新状态
func (n *Tree) checkStatus() {
	_, _, ok := ParseNamedParam(n.path)
	if ok {
		n.nType = PARAM
	} else {
		n.nType = STATIC
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

