package log

import "sync"

type buffer struct {
	bytes        *[]byte
	disableColor bool
}

var bufPool = sync.Pool{
	New: func() any {
		buf := make([]byte, 0, 1024)
		return &buffer{
			bytes: &buf,
		}
	},
}

func newBuffer(handler *handler) *buffer {
	buf := bufPool.Get().(*buffer)
	buf.disableColor = handler.opts.DisableColor
	return buf
}

func (b *buffer) Free() {
	const maxBufferSize = 16 << 10
	if cap(*b.bytes) <= maxBufferSize {
		*b.bytes = (*b.bytes)[:0]
		b.disableColor = false
		bufPool.Put(b)
	}
}
func (b *buffer) Write(bytes []byte) *buffer {
	*b.bytes = append(*b.bytes, bytes...)
	return b
}

func (b *buffer) WriteBytes(char ...byte) *buffer {
	*b.bytes = append(*b.bytes, char...)
	return b
}

func (b *buffer) WriteString(str string) *buffer {
	*b.bytes = append(*b.bytes, str...)
	return b
}

func (b *buffer) WriteColorString(str ...string) *buffer {
	for _, s := range str {
		_, isColor := colors[s]
		if !isColor || !b.disableColor {
			b.WriteString(s)
		}
	}
	return b
}
