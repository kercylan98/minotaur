package mailbox

import "github.com/kercylan98/minotaur/engine/prc"

// Recipient 邮箱接收者接口
type Recipient interface {
	// ProcessUserMessage 处理用户消息
	ProcessUserMessage(message prc.Message)

	// ProcessSystemMessage 处理系统消息
	ProcessSystemMessage(message prc.Message)

	// ProcessAccident 处理意外情况
	ProcessAccident(reason prc.Message)
}
