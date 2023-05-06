package builtin

import (
	"github.com/kercylan98/minotaur/game"
	"github.com/kercylan98/minotaur/utils/huge"
	"github.com/kercylan98/minotaur/utils/synchronization"
)

func NewAttrs() *Attrs {
	return &Attrs{
		attrs: synchronization.NewMap[int, any](),
	}
}

type Attrs struct {
	attrs *synchronization.Map[int, any]

	attrChangeEventHandles   []game.AttrChangeEventHandle
	attrIdChangeEventHandles map[int][]game.AttrChangeEventHandle
}

func (slf *Attrs) SetAttrInt(id int, value int) {
	slf.attrs.Set(id, value)
	slf.OnAttrChangeEvent(id)
	slf.OnAttrIdChangeEvent(id)
}

func (slf *Attrs) SetAttrInt8(id int, value int8) {
	slf.attrs.Set(id, value)
	slf.OnAttrChangeEvent(id)
	slf.OnAttrIdChangeEvent(id)
}

func (slf *Attrs) SetAttrInt16(id int, value int16) {
	slf.attrs.Set(id, value)
	slf.OnAttrChangeEvent(id)
	slf.OnAttrIdChangeEvent(id)
}

func (slf *Attrs) SetAttrInt32(id int, value int32) {
	slf.attrs.Set(id, value)
	slf.OnAttrChangeEvent(id)
	slf.OnAttrIdChangeEvent(id)
}

func (slf *Attrs) SetAttrInt64(id int, value int64) {
	slf.attrs.Set(id, value)
	slf.OnAttrChangeEvent(id)
	slf.OnAttrIdChangeEvent(id)
}

func (slf *Attrs) SetAttrUint(id int, value uint) {
	slf.attrs.Set(id, value)
	slf.OnAttrChangeEvent(id)
	slf.OnAttrIdChangeEvent(id)
}

func (slf *Attrs) SetAttrUint8(id int, value uint8) {
	slf.attrs.Set(id, value)
	slf.OnAttrChangeEvent(id)
	slf.OnAttrIdChangeEvent(id)
}

func (slf *Attrs) SetAttrUint16(id int, value uint16) {
	slf.attrs.Set(id, value)
	slf.OnAttrChangeEvent(id)
	slf.OnAttrIdChangeEvent(id)
}

func (slf *Attrs) SetAttrUint32(id int, value uint32) {
	slf.attrs.Set(id, value)
	slf.OnAttrChangeEvent(id)
	slf.OnAttrIdChangeEvent(id)
}

func (slf *Attrs) SetAttrUint64(id int, value uint64) {
	slf.attrs.Set(id, value)
	slf.OnAttrChangeEvent(id)
	slf.OnAttrIdChangeEvent(id)
}

func (slf *Attrs) SetAttrFloat32(id int, value float32) {
	slf.attrs.Set(id, value)
	slf.OnAttrChangeEvent(id)
	slf.OnAttrIdChangeEvent(id)
}

func (slf *Attrs) SetAttrFloat64(id int, value float64) {
	slf.attrs.Set(id, value)
	slf.OnAttrChangeEvent(id)
	slf.OnAttrIdChangeEvent(id)
}

func (slf *Attrs) SetAttrHugeInt(id int, value *huge.Int) {
	slf.attrs.Set(id, value)
	slf.OnAttrChangeEvent(id)
	slf.OnAttrIdChangeEvent(id)
}

func (slf *Attrs) GetAttrInt(id int) int {
	v, exist := slf.attrs.GetExist(id)
	if !exist {
		return 0
	}
	switch value := v.(type) {
	case float32:
		return int(value)
	case float64:
		return int(value)
	case *huge.Int:
		return int(value.Int64())
	case uint:
		return int(value)
	case uint8:
		return int(value)
	case uint16:
		return int(value)
	case uint32:
		return int(value)
	case uint64:
		return int(value)
	case int:
		return value
	case int8:
		return int(value)
	case int16:
		return int(value)
	case int32:
		return int(value)
	case int64:
		return int(value)
	}
	return 0
}

func (slf *Attrs) GetAttrInt8(id int) int8 {
	v, exist := slf.attrs.GetExist(id)
	if !exist {
		return 0
	}
	switch value := v.(type) {
	case float32:
		return int8(value)
	case float64:
		return int8(value)
	case *huge.Int:
		return int8(value.Int64())
	case uint:
		return int8(value)
	case uint8:
		return int8(value)
	case uint16:
		return int8(value)
	case uint32:
		return int8(value)
	case uint64:
		return int8(value)
	case int:
		return int8(value)
	case int8:
		return value
	case int16:
		return int8(value)
	case int32:
		return int8(value)
	case int64:
		return int8(value)
	}
	return 0
}

func (slf *Attrs) GetAttrInt16(id int) int16 {
	v, exist := slf.attrs.GetExist(id)
	if !exist {
		return 0
	}
	switch value := v.(type) {
	case float32:
		return int16(value)
	case float64:
		return int16(value)
	case *huge.Int:
		return int16(value.Int64())
	case uint:
		return int16(value)
	case uint8:
		return int16(value)
	case uint16:
		return int16(value)
	case uint32:
		return int16(value)
	case uint64:
		return int16(value)
	case int:
		return int16(value)
	case int8:
		return int16(value)
	case int16:
		return value
	case int32:
		return int16(value)
	case int64:
		return int16(value)
	}
	return 0
}

func (slf *Attrs) GetAttrInt32(id int) int32 {
	v, exist := slf.attrs.GetExist(id)
	if !exist {
		return 0
	}
	switch value := v.(type) {
	case float32:
		return int32(value)
	case float64:
		return int32(value)
	case *huge.Int:
		return int32(value.Int64())
	case uint:
		return int32(value)
	case uint8:
		return int32(value)
	case uint16:
		return int32(value)
	case uint32:
		return int32(value)
	case uint64:
		return int32(value)
	case int:
		return int32(value)
	case int8:
		return int32(value)
	case int16:
		return int32(value)
	case int32:
		return value
	case int64:
		return int32(value)
	}
	return 0
}

func (slf *Attrs) GetAttrInt64(id int) int64 {
	v, exist := slf.attrs.GetExist(id)
	if !exist {
		return 0
	}
	switch value := v.(type) {
	case float32:
		return int64(value)
	case float64:
		return int64(value)
	case *huge.Int:
		return value.Int64()
	case uint:
		return int64(value)
	case uint8:
		return int64(value)
	case uint16:
		return int64(value)
	case uint32:
		return int64(value)
	case uint64:
		return int64(value)
	case int:
		return int64(value)
	case int8:
		return int64(value)
	case int16:
		return int64(value)
	case int32:
		return int64(value)
	case int64:
		return int64(value)
	}
	return 0
}

func (slf *Attrs) GetAttrUint(id int) uint {
	v, exist := slf.attrs.GetExist(id)
	if !exist {
		return 0
	}
	switch value := v.(type) {
	case float32:
		return uint(value)
	case float64:
		return uint(value)
	case *huge.Int:
		return uint(value.Int64())
	case uint:
		return value
	case uint8:
		return uint(value)
	case uint16:
		return uint(value)
	case uint32:
		return uint(value)
	case uint64:
		return uint(value)
	case int:
		return uint(value)
	case int8:
		return uint(value)
	case int16:
		return uint(value)
	case int32:
		return uint(value)
	case int64:
		return uint(value)
	}
	return 0
}

func (slf *Attrs) GetAttrUint8(id int) uint8 {
	v, exist := slf.attrs.GetExist(id)
	if !exist {
		return 0
	}
	switch value := v.(type) {
	case float32:
		return uint8(value)
	case float64:
		return uint8(value)
	case *huge.Int:
		return uint8(value.Int64())
	case uint:
		return uint8(value)
	case uint8:
		return value
	case uint16:
		return uint8(value)
	case uint32:
		return uint8(value)
	case uint64:
		return uint8(value)
	case int:
		return uint8(value)
	case int8:
		return uint8(value)
	case int16:
		return uint8(value)
	case int32:
		return uint8(value)
	case int64:
		return uint8(value)
	}
	return 0
}

func (slf *Attrs) GetAttrUint16(id int) uint16 {
	v, exist := slf.attrs.GetExist(id)
	if !exist {
		return 0
	}
	switch value := v.(type) {
	case float32:
		return uint16(value)
	case float64:
		return uint16(value)
	case *huge.Int:
		return uint16(value.Int64())
	case uint:
		return uint16(value)
	case uint8:
		return uint16(value)
	case uint16:
		return value
	case uint32:
		return uint16(value)
	case uint64:
		return uint16(value)
	case int:
		return uint16(value)
	case int8:
		return uint16(value)
	case int16:
		return uint16(value)
	case int32:
		return uint16(value)
	case int64:
		return uint16(value)
	}
	return 0
}

func (slf *Attrs) GetAttrUint32(id int) uint32 {
	v, exist := slf.attrs.GetExist(id)
	if !exist {
		return 0
	}
	switch value := v.(type) {
	case float32:
		return uint32(value)
	case float64:
		return uint32(value)
	case *huge.Int:
		return uint32(value.Int64())
	case uint:
		return uint32(value)
	case uint8:
		return uint32(value)
	case uint16:
		return uint32(value)
	case uint32:
		return value
	case uint64:
		return uint32(value)
	case int:
		return uint32(value)
	case int8:
		return uint32(value)
	case int16:
		return uint32(value)
	case int32:
		return uint32(value)
	case int64:
		return uint32(value)
	}
	return 0
}

func (slf *Attrs) GetAttrUint64(id int) uint64 {
	v, exist := slf.attrs.GetExist(id)
	if !exist {
		return 0
	}
	switch value := v.(type) {
	case float32:
		return uint64(value)
	case float64:
		return uint64(value)
	case *huge.Int:
		return uint64(value.Int64())
	case uint:
		return uint64(value)
	case uint8:
		return uint64(value)
	case uint16:
		return uint64(value)
	case uint32:
		return uint64(value)
	case uint64:
		return value
	case int:
		return uint64(value)
	case int8:
		return uint64(value)
	case int16:
		return uint64(value)
	case int32:
		return uint64(value)
	case int64:
		return uint64(value)
	}
	return 0
}

func (slf *Attrs) GetAttrFloat32(id int) float32 {
	v, exist := slf.attrs.GetExist(id)
	if !exist {
		return 0
	}
	switch value := v.(type) {
	case float32:
		return value
	case float64:
		return float32(value)
	case *huge.Int:
		return float32(value.Int64())
	case uint:
		return float32(value)
	case uint8:
		return float32(value)
	case uint16:
		return float32(value)
	case uint32:
		return float32(value)
	case uint64:
		return float32(value)
	case int:
		return float32(value)
	case int8:
		return float32(value)
	case int16:
		return float32(value)
	case int32:
		return float32(value)
	case int64:
		return float32(value)
	}
	return 0
}

func (slf *Attrs) GetAttrFloat64(id int) float64 {
	v, exist := slf.attrs.GetExist(id)
	if !exist {
		return 0
	}
	switch value := v.(type) {
	case float32:
		return float64(value)
	case float64:
		return value
	case *huge.Int:
		return float64(value.Int64())
	case uint:
		return float64(value)
	case uint8:
		return float64(value)
	case uint16:
		return float64(value)
	case uint32:
		return float64(value)
	case uint64:
		return float64(value)
	case int:
		return float64(value)
	case int8:
		return float64(value)
	case int16:
		return float64(value)
	case int32:
		return float64(value)
	case int64:
		return float64(value)
	}
	return 0
}

func (slf *Attrs) GetAttrHugeInt(id int) *huge.Int {
	v, exist := slf.attrs.GetExist(id)
	if !exist {
		return huge.IntZero.Copy()
	}
	switch value := v.(type) {
	case float32:
		return huge.NewInt(int64(value))
	case float64:
		return huge.NewInt(int64(value))
	case *huge.Int:
		return value.Copy()
	case uint:
		return huge.NewInt(int64(value))
	case uint8:
		return huge.NewInt(int64(value))
	case uint16:
		return huge.NewInt(int64(value))
	case uint32:
		return huge.NewInt(int64(value))
	case uint64:
		return huge.NewInt(int64(value))
	case int:
		return huge.NewInt(int64(value))
	case int8:
		return huge.NewInt(int64(value))
	case int16:
		return huge.NewInt(int64(value))
	case int32:
		return huge.NewInt(int64(value))
	case int64:
		return huge.NewInt(value)
	}
	return huge.IntZero.Copy()
}

func (slf *Attrs) ChangeAttrInt(id int, value int) {
	slf.SetAttrInt(id, slf.GetAttrInt(id)+value)
}

func (slf *Attrs) ChangeAttrInt8(id int, value int8) {
	slf.SetAttrInt8(id, slf.GetAttrInt8(id)+value)
}

func (slf *Attrs) ChangeAttrInt16(id int, value int16) {
	slf.SetAttrInt16(id, slf.GetAttrInt16(id)+value)
}

func (slf *Attrs) ChangeAttrInt32(id int, value int32) {
	slf.SetAttrInt32(id, slf.GetAttrInt32(id)+value)
}

func (slf *Attrs) ChangeAttrInt64(id int, value int64) {
	slf.SetAttrInt64(id, slf.GetAttrInt64(id)+value)
}

func (slf *Attrs) ChangeAttrUint(id int, value uint) {
	slf.SetAttrUint(id, slf.GetAttrUint(id)+value)
}

func (slf *Attrs) ChangeAttrUint8(id int, value uint8) {
	slf.SetAttrUint8(id, slf.GetAttrUint8(id)+value)
}

func (slf *Attrs) ChangeAttrUint16(id int, value uint16) {
	slf.SetAttrUint16(id, slf.GetAttrUint16(id)+value)
}

func (slf *Attrs) ChangeAttrUint32(id int, value uint32) {
	slf.SetAttrUint32(id, slf.GetAttrUint32(id)+value)
}

func (slf *Attrs) ChangeAttrUint64(id int, value uint64) {
	slf.SetAttrUint64(id, slf.GetAttrUint64(id)+value)
}

func (slf *Attrs) ChangeAttrFloat32(id int, value float32) {
	slf.SetAttrFloat32(id, slf.GetAttrFloat32(id)+value)
}

func (slf *Attrs) ChangeAttrFloat64(id int, value float64) {
	slf.SetAttrFloat64(id, slf.GetAttrFloat64(id)+value)
}

func (slf *Attrs) ChangeAttrHugeInt(id int, value *huge.Int) {
	slf.SetAttrHugeInt(id, slf.GetAttrHugeInt(id).Add(value))
}

func (slf *Attrs) RegAttrChangeEvent(handle game.AttrChangeEventHandle) {
	slf.attrChangeEventHandles = append(slf.attrChangeEventHandles, handle)
}

func (slf *Attrs) OnAttrChangeEvent(id int) {
	for _, handle := range slf.attrChangeEventHandles {
		handle(id, slf)
	}
}

func (slf *Attrs) RegAttrIdChangeEvent(id int, handle game.AttrChangeEventHandle) {
	if slf.attrIdChangeEventHandles == nil {
		slf.attrIdChangeEventHandles = map[int][]game.AttrChangeEventHandle{}
	}
	slf.attrIdChangeEventHandles[id] = append(slf.attrIdChangeEventHandles[id], handle)
}

func (slf *Attrs) OnAttrIdChangeEvent(id int) {
	for _, handle := range slf.attrIdChangeEventHandles[id] {
		handle(id, slf)
	}
}
