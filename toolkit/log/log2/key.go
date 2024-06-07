package log

type AttrKey uint16 // 属性键

const (
	AttrKeyTime    AttrKey = iota + 1 // 时间
	AttrKeyLevel                      // 级别
	AttrKeyCaller                     // 调用者
	AttrKeyMessage                    // 消息
)
