package ecs

import (
	"reflect"
	"unsafe"
)

type ComponentId = uint64

func newComponent(typ reflect.Type) *component {
	c := &component{
		typ:      typ,
		size:     typ.Size(),
		align:    uintptr(typ.Align()),
		length:   0,
		capacity: 0,
	}

	return c
}

type component struct {
	typ      reflect.Type   // 组件类型
	size     uintptr        // 组件大小
	align    uintptr        // 对齐方式
	pointer  unsafe.Pointer // 指向组件数据切片的指针
	length   int            // 数据数量
	capacity int            // 数据容量
}

func (c *component) Append(num int) {
	if num <= 0 {
		return
	}
	c.Expand(num)
	c.length += num
}

func (c *component) Set(index int, data unsafe.Pointer) {
	if index < 0 || index >= c.length {
		return
	}

	target := unsafe.Pointer(uintptr(c.pointer) + uintptr(index)*c.size)
	copy((*[1 << 30]byte)(target)[:c.size], (*[1 << 30]byte)(data)[:c.size])
}

func (c *component) Get(index int) unsafe.Pointer {
	if index < 0 || index >= c.length {
		return nil
	}

	return unsafe.Pointer(uintptr(c.pointer) + uintptr(index)*c.size)
}

func (c *component) DeleteRange(start, end int) {
	if start < 0 || start >= c.length || end < 0 || end > c.length || start >= end {
		return
	}

	if end < c.length {
		left := uintptr(start) * c.size
		right := uintptr(end) * c.size
		newSlice := make([]byte, c.size*(uintptr(c.length)-uintptr(end-start)))
		newPointer := unsafe.Pointer(&newSlice[0])
		oldSlice := (*[1 << 30]byte)(c.pointer)
		copy(newSlice[:left], oldSlice[:left])
		copy(newSlice[left:], oldSlice[right:])
		c.pointer = newPointer
	}

	c.length -= end - start
}

func (c *component) Delete(index int) {
	c.DeleteRange(index, index+1)
}

func (c *component) Expand(num int) {
	nextLength := c.length + num
	if nextLength <= c.capacity {
		return
	}

	newCapacity := c.capacity
	if newCapacity == 0 {
		newCapacity = 1
	}
	for newCapacity < nextLength {
		newCapacity *= 2
	}

	newSize := uintptr(newCapacity) * c.size
	offset := (c.align - (newSize % c.align)) % c.align
	newTotalSize := newSize + offset
	newSlice := make([]byte, newTotalSize)
	newPointer := unsafe.Pointer(&newSlice[0])
	if c.pointer != nil {
		oldSlice := (*[1 << 30]byte)(c.pointer)
		copy(newSlice[offset:], oldSlice[:uintptr(c.length)*c.size])
	}

	c.pointer = newPointer
	c.capacity = newCapacity
}

func (c *component) Len() int {
	return c.length
}

func (c *component) Cap() int {
	return c.capacity
}
