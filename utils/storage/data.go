package storage

// newData 创建一个数据存储
func newData[Body any](body Body) *Data[Body] {
	data := &Data[Body]{
		body: body,
	}
	return data
}

// Data 数据存储
//   - 数据存储默认拥有一个 Body 字段
type Data[Body any] struct {
	body Body
}

// Handle 处理数据
func (slf *Data[Body]) Handle(handler func(data Body)) {
	handler(slf.body)
}
