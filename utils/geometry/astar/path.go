package astar

import "github.com/kercylan98/minotaur/utils/generic"

type path[Node comparable, V generic.SignedNumber] []Node

// 获取末位元素
func (p path[Node, V]) last() Node {
	return p[len(p)-1]
}

// 将 n 追加到末位，返回新的路径，不影响原始路径
func (p path[Node, V]) cont(n Node) path[Node, V] {
	cp := make([]Node, len(p), len(p)+1)
	copy(cp, p)
	cp = append(cp, n)
	return cp
}

// cost 计算路径总消耗
func (p path[Node, V]) cost(d func(a, b Node) V) (c V) {
	for i := 1; i < len(p); i++ {
		c += d(p[i-1], p[i])
	}
	return c
}
