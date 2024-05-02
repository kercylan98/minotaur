package server

const (
	connDescriptorInvalidReadDeadline  ConnDescriptor = 1 << iota // 无效的读取截止时间
	connDescriptorInvalidWriteDeadline                            // 无效的写入截止时间
)

// ConnDescriptor 实现 BitSet 实现的连接描述符，用于标识连接部分信息
type ConnDescriptor uint64

// has 判断是否包含指定的标识
func (c *ConnDescriptor) has(flag ConnDescriptor) bool {
	return *c&flag != 0
}

// set 设置指定的标识
func (c *ConnDescriptor) set(flag ConnDescriptor) {
	*c |= flag
}

// SetInvalidReadDeadline 设置无效的读取截止时间
func (c *ConnDescriptor) SetInvalidReadDeadline() {
	c.set(connDescriptorInvalidReadDeadline)
}

// SetInvalidWriteDeadline 设置无效的写入截止时间
func (c *ConnDescriptor) SetInvalidWriteDeadline() {
	c.set(connDescriptorInvalidWriteDeadline)
}
