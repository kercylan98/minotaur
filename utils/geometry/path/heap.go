package path

type h []*Node

func (slf *h) Len() int {
	return len(*slf)
}

func (slf *h) Less(i, j int) bool {
	return (*slf)[i].cost < (*slf)[j].cost
}

func (slf *h) Swap(i, j int) {
	(*slf)[i], (*slf)[j] = (*slf)[j], (*slf)[i]
}

func (slf *h) Push(x any) {
	*slf = append(*slf, x.(*Node))
}

func (slf *h) Pop() any {
	old := *slf
	n := len(old)
	x := old[n-1]
	*slf = old[0 : n-1]
	return x
}
