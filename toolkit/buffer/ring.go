package buffer

// NewRing 创建一个并发不安全的环形缓冲区
//   - initSize: 初始容量
//
// 当初始容量小于 2 或未设置时，将会使用默认容量 2
func NewRing[T any](initSize ...int) *Ring[T] {
	if len(initSize) == 0 {
		initSize = append(initSize, 2)
	}
	if initSize[0] < 2 {
		initSize[0] = 2
	}

	return &Ring[T]{
		buf:      make([]T, initSize[0]),
		initSize: initSize[0],
		size:     initSize[0],
	}
}

// Ring 环形缓冲区
type Ring[T any] struct {
	buf      []T
	initSize int
	size     int
	r        int
	w        int
}

// Read 读取数据
func (b *Ring[T]) Read() (T, error) {
	var t T
	if b.r == b.w {
		return t, ErrBufferIsEmpty
	}

	v := b.buf[b.r]
	b.r++
	if b.r == b.size {
		b.r = 0
	}

	return v, nil
}

// ReadAll 读取所有数据
func (b *Ring[T]) ReadAll() []T {
	if b.r == b.w {
		return nil // 没有数据时返回空切片
	}

	var length int
	var data []T

	if b.w > b.r {
		length = b.w - b.r
	} else {
		length = len(b.buf) - b.r + b.w
	}
	data = make([]T, length) // 预分配空间

	if b.w > b.r {
		copy(data, b.buf[b.r:b.w])
	} else {
		copied := copy(data, b.buf[b.r:])
		copy(data[copied:], b.buf[:b.w])
	}

	b.r = 0
	b.w = 0

	return data
}

// Peek 查看数据
func (b *Ring[T]) Peek() (t T, err error) {
	if b.r == b.w {
		return t, ErrBufferIsEmpty
	}

	return b.buf[b.r], nil
}

// Write 写入数据
func (b *Ring[T]) Write(v T) {
	b.buf[b.w] = v
	b.w++

	if b.w == b.size {
		b.w = 0
	}

	if b.w == b.r {
		b.grow()
	}
}

// grow 扩容
func (b *Ring[T]) grow() {
	var size int
	if b.size < 1024 {
		size = b.size * 2
	} else {
		size = b.size + b.size/4
	}

	buf := make([]T, size)

	copy(buf[0:], b.buf[b.r:])
	copy(buf[b.size-b.r:], b.buf[0:b.r])

	b.r = 0
	b.w = b.size
	b.size = size
	b.buf = buf
}

// IsEmpty 是否为空
func (b *Ring[T]) IsEmpty() bool {
	return b.r == b.w
}

// Cap 返回缓冲区容量
func (b *Ring[T]) Cap() int {
	return b.size
}

// Len 返回缓冲区长度
func (b *Ring[T]) Len() int {
	if b.r == b.w {
		return 0
	}

	if b.w > b.r {
		return b.w - b.r
	}

	return b.size - b.r + b.w
}

// Reset 重置缓冲区
func (b *Ring[T]) Reset() {
	b.r = 0
	b.w = 0
	b.size = b.initSize
	b.buf = make([]T, b.initSize)
}
