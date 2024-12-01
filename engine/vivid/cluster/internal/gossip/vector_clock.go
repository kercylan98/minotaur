package gossip

// newVectorClock 创建一个新的向量时钟
func newVectorClock() *VectorClock {
	return &VectorClock{
		Version: make(map[string]uint64),
	}
}
