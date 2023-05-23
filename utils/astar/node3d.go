package astar

type Node3D struct {
	point   Point3D
	parent  *Node3D
	g, h, f float64
	index   int
}

type Node3DList []*Node3D

func (slf *Node3DList) Len() int {
	return len(*slf)
}

func (slf *Node3DList) Less(i, j int) bool {
	nl := *slf
	return nl[i].f < nl[j].f
}

func (slf *Node3DList) Swap(i, j int) {
	nl := *slf
	nl[i], nl[j] = nl[j], nl[i]
	nl[i].index = i
	nl[j].index = j
}

func (slf *Node3DList) Push(x interface{}) {
	n := len(*slf)
	node := x.(*Node3D)
	node.index = n
	*slf = append(*slf, node)
}

func (slf *Node3DList) Pop() interface{} {
	old := *slf
	n := len(old)
	node := old[n-1]
	node.index = -1
	*slf = old[0 : n-1]
	return node
}
