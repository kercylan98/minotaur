package random

// WeightSlice 按权重随机从切片中产生一个数据并返回
func WeightSlice[T any](getWeightHandle func(data T) int64, data ...T) T {
	item, _ := WeightSliceIndex(getWeightHandle, data...)
	return item
}

// WeightSliceIndex 按权重随机从切片中产生一个数据并返回数据和对应索引
func WeightSliceIndex[T any](getWeightHandle func(data T) int64, data ...T) (item T, index int) {
	var total int64
	var overlayWeight []int64
	for _, d := range data {
		total += getWeightHandle(d)
		overlayWeight = append(overlayWeight, total)
	}
	var r = Int(0, total)
	var i, count = 0, len(overlayWeight)
	for i < count {
		h := int(uint(i+count) >> 1)
		if overlayWeight[h] < r {
			i = h + 1
		} else {
			count = h
		}
	}
	return data[i], i
}

// WeightMap 按权重随机从map中产生一个数据并返回
func WeightMap[K comparable, T any](getWeightHandle func(data T) int64, data map[K]T) T {
	item, _ := WeightMapKey(getWeightHandle, data)
	return item
}

// WeightMapKey 按权重随机从map中产生一个数据并返回数据和对应 key
func WeightMapKey[K comparable, T any](getWeightHandle func(data T) int64, data map[K]T) (item T, key K) {
	var total int64
	var overlayWeight []int64
	var dataSlice = make([]T, 0, len(data))
	var dataKeySlice = make([]K, 0, len(data))
	for k, d := range data {
		total += getWeightHandle(d)
		dataSlice = append(dataSlice, d)
		dataKeySlice = append(dataKeySlice, k)
		overlayWeight = append(overlayWeight, total)
	}
	var r = Int(0, total)
	var i, count = 0, len(overlayWeight)
	for i < count {
		h := int(uint(i+count) >> 1)
		if overlayWeight[h] < r {
			i = h + 1
		} else {
			count = h
		}
	}
	return dataSlice[i], dataKeySlice[i]
}
