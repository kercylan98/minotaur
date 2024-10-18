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
//   - VectorClockOrdering_VCO_Before: 当函数返回 Before，表示 vc 是在 that 之前。也就是说，vc 的事件发生在 that 的事件之前。
//   - VectorClockOrdering_VCO_After: 当函数返回 After，表示 vc 是在 that 之后。也就是说，vc 的事件发生在 that 的事件之后。
//   - VectorClockOrdering_VCO_Concurrent: 当函数返回 Concurrent，表示 vc 和 that 是并发的，即两者没有因果关系。
//   - VectorClockOrdering_VCO_Same: 当函数返回 Same，表示 vc 和 that 是相同的，即它们的时间完全一致。
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
