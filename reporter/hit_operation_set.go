package reporter

// HitOperationSet 命中操作集
type HitOperationSet[Data any] struct {
	input Data
	data  Data
}

// GetInput 获取新输入的数据
func (slf *HitOperationSet[Data]) GetInput() Data {
	return slf.input
}

// GetData 获取当前数据
func (slf *HitOperationSet[Data]) GetData() Data {
	return slf.data
}

// SetData 设置当前数据
func (slf *HitOperationSet[Data]) SetData(data Data) {
	slf.data = data
}
