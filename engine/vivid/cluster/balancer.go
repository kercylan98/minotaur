package cluster

type Balancer interface {
	Select(nodes []*Node) *Node
}

type FunctionalBalancer func(nodes []*Node) *Node

func (f FunctionalBalancer) Select(nodes []*Node) *Node {
	return f(nodes)
}
