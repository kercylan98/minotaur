package log

type (
	ColorType uint16 // 日志颜色类型
)

const (
	ColorTypeTime             ColorType = iota + 1 // 时间颜色
	ColorTypeDebugLevel                            // Debug 级别颜色
	ColorTypeInfoLevel                             // Info 级别颜色
	ColorTypeWarnLevel                             // Warn 级别颜色
	ColorTypeErrorLevel                            // Error 级别颜色
	ColorTypeCaller                                // 调用者颜色
	ColorTypeMessage                               // 消息颜色
	ColorTypeAttrKey                               // 属性键颜色
	ColorTypeAttrValue                             // 属性值颜色
	ColorTypeAttrDelimiter                         // 属性分隔符颜色
	ColorTypeAttrErrorKey                          // 错误键颜色
	ColorTypeAttrErrorValue                        // 错误值颜色
	ColorTypeErrorTrack                            // 错误追踪颜色
	ColorTypeErrorTrackHeader                      // 错误追踪头部颜色
)
