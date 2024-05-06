package rpc

type Reader func(dst any)

// ReadTo 将 Reader 读取的数据写入到目标
func (r Reader) ReadTo(dst any) {
	r(dst)
}
