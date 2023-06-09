package game

import "github.com/kercylan98/minotaur/utils/huge"

// Attrs 游戏属性接口定义
//   - 属性通常为直接读取配置，是否合理暂不清晰，目前不推荐使用
type Attrs interface {
	SetAttrInt(id int, value int)
	SetAttrInt8(id int, value int8)
	SetAttrInt16(id int, value int16)
	SetAttrInt32(id int, value int32)
	SetAttrInt64(id int, value int64)
	SetAttrUint(id int, value uint)
	SetAttrUint8(id int, value uint8)
	SetAttrUint16(id int, value uint16)
	SetAttrUint32(id int, value uint32)
	SetAttrUint64(id int, value uint64)
	SetAttrFloat32(id int, value float32)
	SetAttrFloat64(id int, value float64)
	SetAttrHugeInt(id int, value *huge.Int)

	GetAttrInt(id int) int
	GetAttrInt8(id int) int8
	GetAttrInt16(id int) int16
	GetAttrInt32(id int) int32
	GetAttrInt64(id int) int64
	GetAttrUint(id int) uint
	GetAttrUint8(id int) uint8
	GetAttrUint16(id int) uint16
	GetAttrUint32(id int) uint32
	GetAttrUint64(id int) uint64
	GetAttrFloat32(id int) float32
	GetAttrFloat64(id int) float64
	GetAttrHugeInt(id int) *huge.Int

	ChangeAttrInt(id int, value int)
	ChangeAttrInt8(id int, value int8)
	ChangeAttrInt16(id int, value int16)
	ChangeAttrInt32(id int, value int32)
	ChangeAttrInt64(id int, value int64)
	ChangeAttrUint(id int, value uint)
	ChangeAttrUint8(id int, value uint8)
	ChangeAttrUint16(id int, value uint16)
	ChangeAttrUint32(id int, value uint32)
	ChangeAttrUint64(id int, value uint64)
	ChangeAttrFloat32(id int, value float32)
	ChangeAttrFloat64(id int, value float64)
	ChangeAttrHugeInt(id int, value *huge.Int)

	// RegAttrChangeEvent 任一属性发生变化将立即执行被注册的事件处理函数
	RegAttrChangeEvent(handle AttrChangeEventHandle)
	OnAttrChangeEvent(id int)
	// RegAttrIdChangeEvent 特定属性发生变化将立即执行被注册的事件处理函数
	RegAttrIdChangeEvent(id int, handle AttrChangeEventHandle)
	OnAttrIdChangeEvent(id int)
}

type (
	AttrChangeEventHandle func(id int, attrs Attrs)
)
