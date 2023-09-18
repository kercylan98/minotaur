package buffer

// NewRing 创建一个环形缓冲区
func NewRing[T any](initSize int) *Ring[T] {
	if initSize <= 1 {
		panic("initial size must be great than one")
	}

	return &Ring[T]{
		buf:      make([]T, initSize),
		initSize: initSize,
		size:     initSize,
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
func (slf *Ring[T]) Read() (T, error) {
	var t T
	if slf.r == slf.w {
		return t, ErrBufferIsEmpty
	}

	v := slf.buf[slf.r]
	slf.r++
	if slf.r == slf.size {
		slf.r = 0
	}

	return v, nil
}

// Peek 查看数据
func (slf *Ring[T]) Peek() (t T, err error) {
	if slf.r == slf.w {
		return t, ErrBufferIsEmpty
	}

	return slf.buf[slf.r], nil
}

// Write 写入数据
func (slf *Ring[T]) Write(v T) {
	slf.buf[slf.w] = v
	slf.w++

	if slf.w == slf.size {
		slf.w = 0
	}

	if slf.w == slf.r {
		slf.grow()
	}
}

// grow 扩容
func (slf *Ring[T]) grow() {
	var size int
	if slf.size < 1024 {
		size = slf.size * 2
	} else {
		size = slf.size + slf.size/4
	}

	buf := make([]T, size)

	copy(buf[0:], slf.buf[slf.r:])
	copy(buf[slf.size-slf.r:], slf.buf[0:slf.r])

	slf.r = 0
	slf.w = slf.size
	slf.size = size
	slf.buf = buf
}

// IsEmpty 是否为空
func (slf *Ring[T]) IsEmpty() bool {
	return slf.r == slf.w
}

// Cap 返回缓冲区容量
func (slf *Ring[T]) Cap() int {
	return slf.size
}

// Len 返回缓冲区长度
func (slf *Ring[T]) Len() int {
	if slf.r == slf.w {
		return 0
	}

	if slf.w > slf.r {
		return slf.w - slf.r
	}

	return slf.size - slf.r + slf.w
}

// Reset 重置缓冲区
func (slf *Ring[T]) Reset() {
	slf.r = 0
	slf.w = 0
	slf.size = slf.initSize
	slf.buf = make([]T, slf.initSize)
}
