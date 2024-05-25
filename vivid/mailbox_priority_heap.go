package vivid

type priorityHeap []MessageContext

func (h *priorityHeap) Len() int {
	return len(*h)
}

func (h *priorityHeap) Less(i, j int) bool {
	return (*h)[i].GetPriority() < (*h)[j].GetPriority()
}

func (h *priorityHeap) Swap(i, j int) {
	(*h)[i], (*h)[j] = (*h)[j], (*h)[i]
}

func (h *priorityHeap) Push(x any) {
	*h = append(*h, x.(MessageContext))
}

func (h *priorityHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
