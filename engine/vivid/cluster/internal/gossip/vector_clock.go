package gossip

// newVectorClock 创建一个新的向量时钟
func newVectorClock() *VectorClock {
	return &VectorClock{
		Version: make(map[string]uint64),
	}
}

func (vc *VectorClock) Increment(node string) *VectorClock {
	vc.Version[node]++
	return vc
}

func (vc *VectorClock) Merge(that *VectorClock) *VectorClock {
	for node, time := range that.Version {
		if existingTime, exists := vc.Version[node]; !exists || time > existingTime {
			vc.Version[node] = time
		}
	}
	return vc
}

// CompareTo 比较两个向量时钟，返回事件顺序
func (vc *VectorClock) CompareTo(that *VectorClock) VectorClockOrdering {
	hasBefore, hasAfter := false, false

	for node, time := range vc.Version {
		thatTime, exists := that.Version[node]
		if !exists {
			hasAfter = true // 另一个时钟没有这个节点
		} else if time < thatTime {
			hasBefore = true
		} else if time > thatTime {
			hasAfter = true
		}
	}

	for node, time := range that.Version {
		if _, exists := vc.Version[node]; !exists && time > 0 {
			hasBefore = true // 当前时钟没有这个节点
		}
	}

	if hasBefore && hasAfter {
		return VectorClockOrdering_VCO_Concurrent
	}
	if hasAfter {
		return VectorClockOrdering_VCO_After
	}
	if hasBefore {
		return VectorClockOrdering_VCO_Before
	}
	return VectorClockOrdering_VCO_Same
}

// Prune 移除某个节点的版本信息
func (vc *VectorClock) Prune(node string) *VectorClock {
	if _, exists := vc.Version[node]; exists {
		delete(vc.Version, node)
	}
	return vc
}
