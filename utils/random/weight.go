package random

// WeightSlice 按权重随机从切片中产生一个数据并返回
func WeightSlice[T any](getWeightHandle func(data T) int, data ...T) T {
	var total int
	var overlayWeight []int
	for _, d := range data {
		total += getWeightHandle(d)
		overlayWeight = append(overlayWeight, total)
	}
	var r = IntN(total)
	var i, count = 0, len(overlayWeight)
	for i < count {
		h := int(uint(i+count) >> 1)
		if overlayWeight[h] < r {
			i = h + 1
		} else {
			count = h
		}
	}
	return data[i]
}

// WeightMap 按权重随机从map中产生一个数据并返回
func WeightMap[K comparable, T any](getWeightHandle func(data T) int, data map[K]T) T {
	var total int
	var overlayWeight []int
	var dataSlice = make([]T, 0, len(data))
	for _, d := range data {
		total += getWeightHandle(d)
		dataSlice = append(dataSlice, d)
		overlayWeight = append(overlayWeight, total)
	}
	var r = IntN(total)
	var i, count = 0, len(overlayWeight)
	for i < count {
		h := int(uint(i+count) >> 1)
		if overlayWeight[h] < r {
			i = h + 1
		} else {
			count = h
		}
	}
	return dataSlice[i]
}
